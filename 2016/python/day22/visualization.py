#!/usr/bin/python

disk = [['.'] * 37 for row in range(25)]


def add(line):
    line = line.replace('T', '').split()
    _, x, y = line[0].replace('x', '').replace('y', '').split('-')

    if line[2] == '0':
        node = '_'
    elif int(line[1]) < 100:
        node = '.'
    else:
        node = '#'

    disk[int(y)][int(x)] = node


with open('input.txt') as file_handler:
    map(add, file_handler)

for row in disk:
    print ''.join(row)
