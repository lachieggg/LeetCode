/*
 * COMP20007 Design of Algorithms
 * Semester 1 2016
 *
 * Lachlan Grant
 * March 2016
 *
 * This module implements basic functions related to graphs
 *
 */
#include <stdlib.h>

#include <assert.h>

#include "graph.h"

/* Returns a pointer to a new graph with order vertices */
Graph new_graph(int order) {
  /* define the graph */
  Graph g;
  /* malloc g */
  g = (Graph) malloc(sizeof(struct graph_t));
  assert(g);
  /* number of vertices */
  g -> order = order;
  /* number of edges */
  g -> size = 0;
  /* allocate g->order memory segments, each the size of vertex_t */
  g -> vertices = (Vertex) malloc(g -> order * sizeof(struct vertex_t));
  assert(g -> vertices);
  return g;
}

/* Returns whether aim and vertex are pointing to the same location */
bool ptr_eq(void * aim, void * vertex) {
  if (aim == vertex) {
    return 1;
  }
  return 0;
}

/* Returns whether aim and vertex have the same id */
bool id_eq(void * aim, void * vertex) {
  if (((Vertex)(aim)) -> id == ((Vertex) vertex) -> id) {
    return 1;
  }
  return 0;
}

/* Add the edge from v1 to v2 to graph */
void add_edge(Graph graph, int v1, int v2) {

}

/* Free the memory allocated to the whole graph*/
void free_graph(Graph graph) {
  int i;
  if (graph) {
    for (i = 0; i < graph -> order; i++) {
      free(graph -> vertices[i].label);
      if (graph -> vertices[i].in) {
        free_list(graph -> vertices[i].in);
      }
      if (graph -> vertices[i].out) {
        free_list(graph -> vertices[i].out);
      }
    }
    free(graph);
  }
}