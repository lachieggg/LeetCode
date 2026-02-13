#!/bin/bash

gcc mergeSort.c
time ./a.out
time python3 mergeSort.py
rm a.out
