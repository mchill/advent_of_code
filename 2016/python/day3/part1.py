#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.readlines()

possible = 0

for line in content:
    sides = [int(side) for side in line.split()]

    if max(sides) < sum(sides) / 2.0:
        possible += 1

print possible
