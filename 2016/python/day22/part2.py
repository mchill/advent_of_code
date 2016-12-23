#!/usr/bin/python

def construct(line):
    line = line.translate(None, 'Txy').split()
    _, x, y = line[0].split('-')
    return {'pos': (int(x), int(y)), 'size': int(line[1]), 'used': line[2]}


with open('input.txt') as file_handler:
    nodes = [construct(line) for line in file_handler]

empty = next(node for node in nodes if node['used'] == '0')
wall = next(node for node in nodes if node['size'] > 500)

print empty['pos'][1] + 2 * (empty['pos'][0] - wall['pos'][0] + 1) + 176
