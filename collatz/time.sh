#!/bin/bash

gcc collatz.c
printf "C time\n";
time ./a.out 500000;
printf '\n\nPython time\n';
time ./collatz.py 500000;
