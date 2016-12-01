#!/usr/bin/python

class Found(Exception): pass

with open('input.txt') as file_handler:
    content = file_handler.read()

position = (0, 0)
direction = 0
directions = [(0, -1), (1, 0), (0, 1), (-1, 0)]
visited = [(0, 0)]

try:
    for instruction in content.rstrip().split(', '):
        direction = (direction + (1 if instruction[:1] == 'R' else -1)) % 4
        for step in range(0, int(instruction[1:])):
            position = tuple(map(sum, zip(position, directions[direction])))
            if position in visited: raise Found
            visited.append(position)
except Found:
    print sum(map(abs, position))
