#!/bin/bash

# Rust
rustc -A warnings collatz.rs
./collatz 500001

# C
gcc -o collatz collatz.c
./collatz 500001

# Python
python collatz.py 500001

# Cleanup
rm collatz
