#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

#define LOGGING false

bool collatz(int n) {
	if(LOGGING) { printf("n = %d\n", n); }

	if(n == 1) {
		printf("Complete.\n");
		return true;
	}

	if(n % 2 == 0) { // even
		return collatz(n/2);
	} else { // odd
		return collatz(3*n+1);
	}
}

int main(int argc, char **argv) {
	if(argc != 2) {
		return 1; // error
	}

	int n = atoi(argv[1]);
	if(LOGGING) {printf("%d", n); }
	collatz(n);
	return 0;
}

void scan() {
	int n;
	scanf("%d", &n);
	collatz(n);
}
