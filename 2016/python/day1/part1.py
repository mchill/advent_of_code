#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.read()

position = (0, 0)
direction = 0
directions = [(0, -1), (1, 0), (0, 1), (-1, 0)]

for instruction in content.rstrip().split(', '):
    direction = (direction + (1 if instruction[:1] == 'R' else -1)) % 4
    position = tuple(map(sum, zip(position, tuple(
        [int(instruction[1:]) * val for val in directions[direction]]
    ))))

print sum(map(abs, position))
