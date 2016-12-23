#!/usr/bin/python
import itertools


def construct(line):
    return {'used': int(line[2]), 'avail': int(line[3])}


with open('input.txt') as file_handler:
    nodes = [construct(line.replace('T', '').split()) for line in file_handler]

combos = list(itertools.combinations(nodes, 2))
combos += [reversed(combo) for combo in combos]

print sum(0 < node1['used'] <= node2['avail'] for node1, node2 in combos)
