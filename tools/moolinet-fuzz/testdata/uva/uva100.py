#!/usr/bin/python3
import sys

cache = {}
def algo(n):
    if n not in cache:
        if n == 1: cache[n] = 1
        elif n % 2 == 1: cache[n] = 1 + algo(3*n+1)
        else: cache[n] = 1 + algo(n/2)
    return cache[n]

for line in sys.stdin:
    n, m = [int(n) for n in line.split()]
    res = 0
    for i in range(min(n,m), max(n+1,m+1), 1):
        res = max(res, algo(i))

    print(n, m, res)
