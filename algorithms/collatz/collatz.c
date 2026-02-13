#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

bool collatz(int n) {
	printf("%d ", n);

	if (n == 1) {
		return true;
	}

	if (n % 2 == 0) {
		return collatz(n/2);
	} 
	return collatz(3*n+1);
}

int main(int argc, char **argv) {
	int n = 101010101;
	if (argc == 2) {
		n = atoi(argv[1]);
	}

	collatz(n);
}
