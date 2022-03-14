/*
 * COMP20007 Design of Algorithms
 * Semester 1 2016
 *
 * Lachlan Grant
 * March 2016
 *
 * This module provides all the IO functionality related to graphs.
 *
*/

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>

#include "graphio.h"

#define MAX_LINE_LEN 256
#define INITIAL_SIZE 32
#define EDGE_ERROR -1


int get_nth_int(char *string, int n);
void init_edges(Graph g, FILE *fp);
void init_vertices(Graph g, FILE *fp);

/* Loads the graph from input */
Graph load_graph(char *input) {
	/* For reading the file into */
	char line[MAX_LINE_LEN];

	/* Define and open the file */
	FILE *fp;
	fp = (FILE*)fopen(input, "r");
	fgets(line, MAX_LINE_LEN, fp);

	/* Define the Graph, g */
	Graph g;
	g = new_graph(atoi(line));

	/* Initialize the vertices and edges*/
	init_vertices(g, fp);
	init_edges(g, fp);

	/* Close and return */
	fclose(fp);
	return g;
}

/* Initializes and labels all the vertices */
void init_vertices(Graph g, FILE *fp) {
	int i;
	for(i=0; i<g->order; i+=1) {
		g->vertices[i].id = i;
		/* Allocate memory and assert the components of the vertex label */
		g->vertices[i].label = (char *)malloc(MAX_LINE_LEN*sizeof(char));
		assert(g->vertices[i].label);

		/* Initialize as NULL, then a push will create a new list */
		g->vertices[i].in = NULL;
		g->vertices[i].out = NULL;

		/* Store the line as the label */
		fgets(g->vertices[i].label, MAX_LINE_LEN, fp);
	}
}

/* Initializes and pushes all of the edges into in/out lists */
void init_edges(Graph g, FILE *fp) {
	/* Represents the indexes of the two vertices */
	int ti, tj;
	/* Represents a line from the file */
	char line[MAX_LINE_LEN];
	/* Represents a pointer to a list_t, for the in/out lists of ti/tj */
	List list;

	while((fgets(line, MAX_LINE_LEN, fp)) != NULL) {
		/* ti is on the in of tj, and tj is on the out of ti */
		ti = get_nth_int(line, 0); /* id of ti */
		tj = get_nth_int(line, 1); /* id of tj */

		/* Push (vertex tj) to the in list of ti */
		list = g->vertices[ti].in;
		g->vertices[ti].in = push(list, (void *)&g->vertices[tj]);

		/* Push (vertex ti) to the out list of tj */
		list = g->vertices[tj].out;
		g->vertices[tj].out = push(list, (void*)&g->vertices[ti]);

		/* increment the size of g */
		g->size += 1;
	}
}


/* Returns the nth integer in a string */
int get_nth_int(char *string, int n) {
	/* The index of the word we want is 'n' */
	/* The beginning index of the int we want is 'start' */
	int start = 0;
	while(n>0) {
		if(string[start] == '\0') {
			fprintf(stderr, "Error: Invalid input file\n");
			exit(EXIT_FAILURE);
		} else if(string[start] == ' ') {
			n-=1;
		}
		start+=1;
	}
	return atoi(string+start);
}


/* Prints the graph */
void print_graph(char *output, Graph graph) {
	int i;
	List list;
	FILE *fp;
	fp = (FILE*)fopen(output, "w");

	fputs("digraph {\n", fp);
	for(i=0; i<graph->order; i++) {
		/* The incoming vertices */
		list = graph->vertices[i].in;
		fputs(" ", fp);
		/* Print the label of the current vertex */
		print_vertex_label(fp, (void*)&graph->vertices[i]);
		if(list != NULL) {
			fputs(" -> {", fp);
			/* Print all of the vertices in in, until end */
			while(list != NULL) {
				print_vertex_label(fp, list->data);
				fputs(" ", fp);
				list = list->next;
			}
			fputs("}\n", fp);
		} else {
			fputs("\n", fp);
		}
	}
	fputs("}\n", fp);
}

/* Prints the destination vertex label after a space */
void print_vertex_label(FILE *file, void *vertex) {
	Vertex v;
	char label[MAX_LINE_LEN];
	if (vertex) {
		/* Modify a separate variable */
		v = (Vertex)vertex;
		strcpy(label, v->label);
		/* Remove the newline */
		label[strlen(label)-1] = '\0';
		/* Print the label */
		fputs(" ", file);
		fputs(label, file);
	}
}

/* Prints the id of a vertex then a newline */
void print_vertex_id(FILE *file, void *vertex) {
	if (vertex) {
		fprintf(file, "%d\n", ((Vertex)vertex)->id);
	}
}

/* Returns a sequence of vertices read from file */
List load_vertex_sequence(FILE *file, Graph graph) {
    const char *err_duplicate = "Error: duplicate vertex %d %s\n";
    const char *err_order = "Error: graph order %d, loaded %d vertices\n";
    List list = NULL;
    int id;

    while(fscanf(file, "%d\n", &id) == 1) {
        assert(id >= 0);
        assert(id < graph->order);

        if (!insert_if(id_eq, graph->vertices + id, &list)) {
            fprintf(stderr, err_duplicate, id, graph->vertices[id].label);
            exit(EXIT_FAILURE);
        }
    }

    if (len(list) != graph->order) {
        fprintf(stderr, err_order, graph->order, len(list));
        exit(EXIT_FAILURE);
    }

    return list;
}
