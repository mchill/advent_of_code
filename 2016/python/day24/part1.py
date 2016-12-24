#!/usr/bin/python
from collections import deque
from itertools import permutations
from sys import maxint


def valid_moves(current, traversed, map):
    x, y = current
    moves = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]
    return [move for move in moves if valid(move, traversed, map)]


def valid(position, traversed, map):
    if position not in traversed and map[position[1]][position[0]] != '#':
        traversed.add(position)
        return True


with open('input.txt') as file_handler:
    map = [line.rstrip() for line in file_handler]

distances = {}
positions = ['0', '1', '2', '3', '4', '5', '6', '7']

for start in positions:
    for target in positions:
        if start == target:
            continue

        row = next(row for row in map if start in row)
        position = (row.index(start), map.index(row))

        moves = deque([position])
        traversed = {position}

        index = 0

        while not any(map[y][x] == target for x, y in moves):
            index += 1
            old_moves = moves
            moves = []

            for move in old_moves:
                moves.extend(valid_moves(move, traversed, map))

        distances[(start, target)] = index

steps = maxint
positions.pop(0)

for path in permutations(positions):
    distance = distances[('0', path[0])]

    for index in range(len(path) - 1):
        distance += distances[(path[index], path[index + 1])]

    steps = min(steps, distance)

print steps