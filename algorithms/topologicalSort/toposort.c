/*
 * COMP20007 Design of Algorithms
 * Semester 1 2016
 *
 * Lachlan Grant
 * March 2016
 *
 * This module implements topological sorting algorithms
 *
 */
#include <stdlib.h>

#include <stdio.h>

#include <assert.h>

#include "graphio.h"

#include "toposort.h"

#include "graph.h"

#define NOT_VISITED 0
#define VISITED 1
#define VERTEX_NOT_FOUND - 1

List visit(Vertex N, List l, int * temp, int * perm);
List init_starts(Graph graph, List list);
int find_vertex_index(Vertex vertex, List vertices);
int del_node_from_list(Vertex node, List list);

/* A list containing all "start nodes" with no in edges */
List init_starts(Graph graph, List list) {
  /* initialize S */
  int i;
  int length;
  length = 0;
  for (i = 0; i < graph -> order; i++) {
    if (graph -> vertices[i].in == NULL) {
      /* Has no incoming edges */
      length += 1;
      list = push(list, (void * ) & graph -> vertices[i]);
    }
  }
  return list;
}

int del_node_from_list(Vertex node, List list) {
  if (!del(ptr_eq, node, & (list))) {
    fprintf(stderr, "Could not delete vertex\n");
    return 0;
  }
  return 1;
}

/* Returns a list of topologically sorted vertices using the Kahn method */
List kahn_sort(Graph graph) {
  int i;
  printf("Kahn Sort:\n");

  /* Number of graphs with no incoming edges */
  /* The set of vertices which are sorted */
  Vertex N = (struct vertex_t * ) malloc(sizeof(struct vertex_t));
  assert(N);
  Vertex M = (struct vertex_t * ) malloc(sizeof(struct vertex_t));
  assert(M);

  /* The head of the out list of node N */
  List Out = (List) malloc(sizeof(List));

  /* The set of sorted elements */
  List * Sorted = (List * ) malloc(sizeof(List));
  List sorted = (List) push(NULL, NULL);
  * Sorted = sorted;

  /* The set of vertices with no incoming edges */
  List * Starts = (List * ) malloc(sizeof(List));
  List starts = (List) push(NULL, NULL);
  * Starts = starts;

  starts = init_starts(graph, starts);

  /* Produce the Sorted list */
  while (starts != NULL) {
    /* Remove a node N from starts */
    N = starts -> data;
    /* Put N into Sorted */
    insert(N, Sorted);
    /* Slide along s */
    starts = starts -> next;
    /* Give N an iterator */
    Out = N -> out;
    while (Out != NULL) {
      /* M is the next object in Out */
      M = Out -> data;
      if (M -> in == NULL) {
        /* M has no other in nodes */
        break;
      } else if (M -> in -> next != NULL) {
        /* delete the node N from M->in */
        if (!del_node_from_list(N, M -> in )) {
          break;
        }
        /* M has another in node */
        starts = push(starts, M);
      } else {
        if (!(del_node_from_list(N, M -> in ))) {
          break;
        }
      }
      /* Delete M from N->out */
      del_node_from_list(M, N -> out);
      Out = Out -> next;
    }
  }

  for (i = 0; i < graph -> order; i++) {
    if (graph -> vertices[i].in != NULL || graph -> vertices[i].out != NULL) {
      /* The vertex does have edges */
      /* The graph must contain a cycle */
      fprintf(stderr, "The graph contains a cycle, not A DAG\n");
      exit(EXIT_FAILURE);
    }
  }
  return *Sorted;
}

/* Recursively visit all vertices connected to vertex */
List visit(Vertex N, List l, int * temp, int * perm) {
  /* Point to the current element we are exploring  */
  List current = N -> out;

  if (temp[N -> id]) {
    /* the graph contains a cycle */
    fprintf(stderr, "Already visited %d, not a DAG\n", N -> id);
    exit(EXIT_FAILURE);
  }

  if (!(perm[N -> id])) {
    /* The vertex has not been marked */
    /* Mark the vertex as temporarily visited */
    temp[N -> id] = VISITED;

    while (current && current -> data) {
      visit(current -> data, l, temp, perm);
      current = current -> next;
    }
    perm[N -> id] = VISITED;
    temp[N -> id] = NOT_VISITED;
    l = push(l, N);
  }
  return l;
}

/* Returns a list of topologically sorted vertices using the DFS method */
List dfs_sort(Graph graph) {
  int i;

  /* Create two auxiliary arrays */
  int temp[graph -> order];
  int perm[graph -> order];

  /* Create the list of sorted vertices */
  List * head = (List * ) malloc(sizeof(List));
  assert(head);
  List l = (List) push(NULL, NULL);

  /* L permanently stores the address of the head of l */
  * head = l;

  printf("DFS Sort:\n");

  /* Initialize all vertices as not visited */
  for (i = 0; i < graph -> order; i++) {
    temp[i] = NOT_VISITED;
    perm[i] = NOT_VISITED;
  }

  /* Visit all vertices that N is connected to */
  for (i = 0; i < graph -> order; i++) {
    l = visit( & graph -> vertices[i], l, temp, perm);
  }

  return *head;
}

/*
 * Return the index of vertex in the list vertices
 * If not found will return VERTEX_NOT_FOUND
 */
int find_vertex_index(Vertex vertex, List vertices) {
  int i = 0;
  while (vertices != NULL && vertices -> data != NULL) {
    if (id_eq((void * ) & vertices[i], (void * ) vertex)) {
      return i;
    }
    vertices = vertices -> next;
    i++;
  }
  return VERTEX_NOT_FOUND;
}

/* Uses graph to verify vertices are topologically sorted */
bool verify(Graph graph, List vertices) {
  int i, u_index, v_index;
  List Out, In;
  Vertex u, v;
  /* Check that for all u's which have an edge from u to v
   * u comes before v in the ordering
   * Check that for all v's which have an edge from v to u
   * v comes before u in the ordering
   * Check that all vertices in graph->vertices are in vertices
   */
  for (i = 0; i < graph -> order; i++) {
    /* u is a vertex from graph->vertices */
    u = & graph -> vertices[i];
    Out = u -> out;
    In = u -> in ;

    while (Out != NULL) {
      /* v is a vertex which has an edge from (u,v) */
      v = Out -> data;
      u_index = find_vertex_index(u, vertices);
      v_index = find_vertex_index(v, vertices);
      if (u_index == VERTEX_NOT_FOUND || v_index == VERTEX_NOT_FOUND) {
        /* u is not in vertices or v is not in vertices */
        /* Thus the vertices are incomplete */
        return 0;
      }
      if (u_index > v_index) {
        /* u does not come before v in the ordering */
        /* The vertices are not a topological sort of the graph */
        return 0;
      }
      v = In -> data;
      v_index = find_vertex_index(v, vertices);
      if (v_index > u_index) {
        /* v does not come before u in the ordering */
        /* vertices is not a topological sort */
        return 0;
      }
      /* Get the next vertex in In */
      In = In -> next;
      /* Get the next vertex in Out */
      Out = Out -> next;
    }
  }
  /* Vertices is a topological sort for graph */
  return 1;
}