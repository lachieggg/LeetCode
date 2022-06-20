#!/bin/bash

# C 
gcc collatz.c -o collatz.o 
printf "C time\n";
time ./collatz.o 500001;

# Rust
printf '\n\nRust time\n';
rustc -A warnings collatz.rs 
time ./collatz 500001;

# Python
printf '\n\nPython time\n';
time ./collatz.py 500001;

# Cleanup
rm collatz.o;
rm collatz
