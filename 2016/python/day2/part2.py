#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.readlines()

KEYPAD = [[None, None, None, None, None, None, None],
          [None, None, None,  1,   None, None, None],
          [None, None,  2,    3,     4,  None, None],
          [None,  5,    6,    7,     8,    9 , None],
          [None, None, 'A',  'B',   'C', None, None],
          [None, None, None, 'D',  None, None, None],
          [None, None, None, None, None, None, None]]
MOVES = {'U': (0, -1), 'R': (1, 0), 'D': (0, 1), 'L': (-1, 0)}

code = ''
button = (1, 3)

for line in content:
    for move in line.rstrip():
        new_button = tuple(map(sum, zip(button, MOVES[move])))
        if KEYPAD[new_button[1]][new_button[0]] is not None:
            button = new_button

    code += str(KEYPAD[button[1]][button[0]])

print code
