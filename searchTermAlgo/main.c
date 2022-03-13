#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

#define MAX_LINES 10
#define MAX_CHARS 1000
#define MAX_QUERY 100
#define TERM_WIDTH 50
#define RANK_HEIGHT 10
#define VERBOSE 1

/* Authored by Lachlan Grant, September 2015 */

/* This program, given a text file of arbitrary size,
*  searches for the line which resembles the search term
*  as closely as possible. The program will return a rank
*  of the lines in the text which achieved the highest level
*  of similarity.
*
*  Resemblance is defined by the number of consecutively
*  correct characters a line has with respect to a query.
*
*/

/* Definition of the line_t struct */
typedef struct {
	int length;
	int line_index;
	int score;
	char text[MAX_CHARS];
} line_t;

/* Function prototype definitions */
void print_score_array(line_t lines[MAX_LINES], int length);
void swap(line_t *p1, line_t *p2);
void sort_scores(line_t list[MAX_CHARS], int length);
void print_scores(line_t *lines, int line_index, int guard);
void print_border();
int get_score(char *query, char *text);


int main(int argc, char **argv) {
  /*
  'char_index'  - current character index
  'line_index'  - current line index
  'score'   		 - score of the current line
  'lines'       - array of the 10 line_t's currently available in memory
  'rlength'     - the length of the lines array (retained length)
  'c'           - the current character
  */
	int char_index, line_index;
	int rlength;

	int score;
	char c;
	line_t lines[MAX_LINES];

	rlength = 0, char_index = 0, line_index = 0;
	line_t line = {char_index, line_index, 0};

	while((c=getchar()) != EOF) {
		if(VERBOSE) {
			printf("%c", c);
		}
		line.text[char_index] = c; /* add the character to the line */
		char_index+=1;
		if(c == '\n') {
			line.text[char_index] = '\0';

			/*
			Calculate the score of a particular line
			If line longer than query, search line for query subsets
			If query longer than line, search query for line subsets
			*/
			if(char_index >= (int)strlen(argv[1])) {
				score = get_score(argv[1], line.text);
			} else {
				score = get_score(line.text, argv[1]);
			}

			/* assign values to line */
			line.length = char_index;
			line.line_index = line_index;
			line.score = score;

			/* retain the value of line */
			if(rlength == MAX_LINES) { /* full */
				if(line.score > lines[rlength-1].score) {
					lines[rlength-1] = line;
				}
			} else { /* non full */
				/* add the line to the next index */
				lines[rlength] = line;
				rlength += 1;
			}
			/* sort the lines */
			sort_scores(lines, rlength);

			line_index+=1;
			char_index=0;
		}
	}

	print_border();
	printf("\nLength of your Query is %d...\n", (int)strlen(argv[1]));
	printf("Thus %d is the highest possible score\n", (int)strlen(argv[1]));
	print_border();

	if(line_index > RANK_HEIGHT) {
		print_scores(lines, line_index, RANK_HEIGHT);
	} else {
		print_scores(lines, line_index, line_index);
	}
	print_border();
	return 0;
}

void print_border() {
	int i;
	printf("\n");
	for(i=0; i<TERM_WIDTH; i++) {
		printf("-");
	}
	printf("\n");
}

/* Print out any arbitrary array of integers */
void print_score_array(line_t lines[MAX_LINES], int length) {
	int i;
	printf("[");
	for(i=0; i<length; i++) {
		if(i!=length-1) {
			printf("%d, ", lines[i].score);
		} else {
			printf("%d", lines[i].score);
		}
	}
	printf("]\n");
}

void print_scores(line_t *lines, int line_index, int guard) {
	int i, j;
	for(i=0; i<guard; i++) {
		printf("%d: Line %4d, score = %d\n", i+1, lines[i].line_index+1,
		lines[i].score);
		for(j=0; j<lines[i].length; j++) {
			printf("%c", lines[i].text[j]);
		}
		printf("---\n");
	}
}


/* Swaps two line_t variables in memory */
void swap(line_t *p1, line_t *p2) {
	line_t temp;
	temp = *p1;
	*p1 = *p2;
	*p2 = temp;
}

/* Sort the list of line_t's based on score
 * Insertion sort, adjusted for line_t's
 */
void sort_scores(line_t list[MAX_CHARS], int length) {
	int i, j;
	for(i=0; i<length; i++) {
		for(j=i-1; j>=0 && list[j+1].score>list[j].score; j--) {
			swap(&list[j], &list[j+1]);
		}
	}
}

/* Generate a score for a particular line and search query */
int get_score(char *query, char *text) {
	/*
	'index'         - index being checked
	'start'         - starting index in text
	'match_length'  - the length of the current pattern match
	'longest_match' - the current highest pattern match length
	*/
	int index, start, match_length, longest_match;
	start=0, match_length=0, longest_match=0;
	int tlength = strlen(text);
	int qlength = strlen(query);

	while(start < tlength-qlength+1) {
		match_length=0;
		for(index=0; index < qlength; index++) {

			if(*(text+start+index)==*(query+index)) {
				match_length+=1;
			} else {
				match_length = 0;
			}

			if(match_length > longest_match) {
				longest_match = match_length;
			}
		}
		start+=1;
	}
	return longest_match;
}
