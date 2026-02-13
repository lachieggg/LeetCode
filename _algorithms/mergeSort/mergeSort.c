#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <assert.h>
#include <string.h>

#define LENGTH 33
#define MAX_NUMBER_VALUE 10

int cmp(int x1, int x2);
void print_array(int *A, int length);
int *mergeSort(int *A,  int n) ;
int *merge(int *A, int *B, int lengthA, int lengthB);

// Print out an array
void print_array(int A[], int length) {
	if(length == 0) {
		printf("\n");
		return;
	}

	int i;
	printf("[");
	for(i=0; i<length-1; i++) {
		printf("%d, ", A[i]);
	}
	printf("%d]\n", A[length-1]);
	return;
}

// Comparison function
int cmp(int x1, int x2) {
	if(x1 > x2) {
		return 1;
	}
	if (x1 < x2) {
		return -1;
	}
	return 0;
}

// Copy an array X into an array Y begginning with index 'start'
int *copyArray(int X[], int Y[], int start, int end) {
	int i;
	for(i=start; i<end; i++) {
		Y[i] = X[i];
	}
	return Y;
}

// Prepend X to Y
int *prepend(int X, int *Y, int lengthY) {
	int lengthP = lengthY + 1;
	int *P = (int*)malloc((lengthP)*sizeof(int));

	// Add X to the start
	P[0] = X;
	// Add all subsequent numbers
	int i;
	for(i=0; i<lengthY; i++) {
		P[i+1] = Y[i];
	}

	return P;
}

// Merge
// Recursively prepend the lowest value from the 
// start of both arrays
int *merge(int *A, int *B, int lengthA, int lengthB) {
	print_array(A, lengthA);
	print_array(B, lengthB);

	if(lengthA == 0) {
		return B;
	}
	if(lengthB == 0) {
		return A;
	}

	if(A[0] <= B[0]) {
		return prepend(A[0], merge(A+1, B, lengthA-1, lengthB), lengthA+lengthB-1);
	
	} else {
		return prepend(B[0], merge(A, B+1, lengthA, lengthB-1), lengthA+lengthB-1);
	}
}

// Generate an array of random integers
int *generate_random_array(int *A, int length) {
	int i, r;
	for(i=0; i<length; i++) {
		r = rand() % 10;
		A[i] = r;
	}

	return A;
}

// Recursive merge sort
int *mergeSort(int *A, int length) {
	if(length > 1) {
		int lengthA = length/2;
		int lengthB = length/2;

		if(length % 2 != 0) {
			lengthB = lengthA + 1;
		}

		int *B;
		B = A+(lengthA);

		print_array(A, lengthA);
		print_array(B, lengthB);

		A = mergeSort(A, lengthA);
		B = mergeSort(B, lengthB); 

		return merge(A, B, lengthA, lengthB);
	} else {
		return A;
	}
}

// Seed the RNG
void seed() {
	time_t t;
	srand((unsigned) time(&t));
}

// Generate an empty array of particular length
int *make_array(int length) {
	int *A;
	A = (int*)malloc(length*sizeof(int));
	assert(A);
	return A;

}
// Main
int main(int argc, char **argv) {
	int length;
	int *A;

	// Seed the RNG
	seed();

	length = LENGTH;
	// Generate an empty array of particular length
	A = make_array(length);
	// Fill the array with random numbers
	generate_random_array(A, length);
	// Sort the array and then print
	print_array(mergeSort(A, length), length);

	return 0;
}


