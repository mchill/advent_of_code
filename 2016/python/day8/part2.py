#!/usr/bin/python
import re
from copy import deepcopy

with open('input.txt') as file_handler:
    instructions = file_handler.readlines()

screen = [[' ' for width in range(50)] for height in range(6)]

for instruction in instructions:
    instruction = re.split(r' |x(?=\d)|=', instruction.rstrip())

    if instruction[0] == 'rect':
        for row in range(int(instruction[2])):
            for col in range(int(instruction[1])):
                screen[row][col] = '#'
        continue

    index = int(instruction[3])
    offset = int(instruction[5])

    if instruction[1] == 'row':
        screen[index] = screen[index][-offset:] + screen[index][:-offset]
    elif instruction[1] == 'column':
        screen = zip(*screen)
        screen[index] = screen[index][-offset:] + screen[index][:-offset]
        screen = [list(row) for row in zip(*screen)]

print('\n'.join([''.join([pixel for pixel in row]) for row in screen]))
