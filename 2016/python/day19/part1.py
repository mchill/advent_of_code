#!/usr/bin/python

elves = range(3004953)
start = 0

while len(elves) > 1:
    new_start = len(elves) + start & 1
    elves = elves[start::2]
    start = new_start

print elves[0] + 1
