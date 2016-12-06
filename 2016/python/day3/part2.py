#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.readlines()

data = []
possible = 0

for line in content:
    data.append([int(side) for side in line.split()])

for row in xrange(0, len(data), 3):
    for col in range(0,3):
        sides = [data[row][col], data[row + 1][col], data[row + 2][col]]

        if max(sides) < sum(sides) / 2.0:
            possible += 1

print possible
