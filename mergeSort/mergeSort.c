#include <stdio.h>
#include <cstdlib>
#include <time.h>

#define LENGTH 16
#define MAX_NUMBER_VALUE 10

int cmp(int x1, int x2);
int *merge(int A[], int B[], int n);
int *mergeSort(int A[],  int n) ;
void print_array(int A[], int length);

// Print out an array
void print_array(int A[], int length) {
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

// Prepend X to A
int *prepend(int X, int A[], int lengthA) {
	int lengthP = lengthA + 1;
	int *P = (int*)malloc((lengthA+1)*sizeof(int));

	// Add X to the start
	P[0] = X;
	// Add all subsequent numbers
	int i;
	for(i=0; i<lengthA; i++) {
		P[i+1] = A[i];
	}

	print_array(P, lengthP);

	//free(A);
	return P;
}

// Merge
int *merge(int A[], int B[], int M[], int startA, int startB, int lengthA, int lengthB) {
	printf("A: ");
	print_array(A, lengthA);
	printf("B: ");
	print_array(B, lengthB);
	printf("M: ");
	print_array(M, 2*LENGTH);
	printf("\n\n");

	if(lengthA == 0) {
		return B;
	}
	if(lengthB == 0) {
		return A;
	}

	if(A[0] <= B[0]) {
		return prepend(A[0], merge(A, B, M, startA+1, startB, lengthA-1, lengthB), 2*LENGTH);
	} else {
		return prepend(B[0], merge(A, B, M, startA, startB+1, lengthA, lengthB-1), 2*LENGTH);
	}

	return M;
}

// Generate an array of random integers
int *generate_random_array(int *A, int length) {
	int i;
	int r;
	for(i=0; i<LENGTH; i++) {
		r = rand() % 10;
		A[i] = r;
	}

	return A;
}

void seed() {
	time_t t;
	srand((unsigned) time(&t));
}

int main(int argc, char **argv) {
	int length = LENGTH;
	int *A = (int*)malloc(length*sizeof(int));

	seed();

	generate_random_array(A, length);

	print_array(A, length);

	return 0;
}


int *mergeSort(int *A, int n) {
	if(n > 1) {
		int *B = A+(n/2)*sizeof(int);
		mergeSort(A, n/2);
		mergeSort(B, n/2); // assume n is even
		int *M = (int*)malloc(n*sizeof(int));
		return merge(A, B, M, 0, 0, n/2, n/2);
	} else {
		return A;
	}
}
