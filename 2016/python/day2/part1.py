#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.readlines()

code = ''
button = 5

for line in content:
    for move in line.rstrip():
        if move == 'U' and button > 3:
            button -= 3
        elif move == 'R' and button % 3 != 0:
            button += 1
        elif move == 'D' and button < 7:
            button += 3
        elif move == 'L' and button % 3 != 1:
            button -= 1

    code += str(button)

print code
