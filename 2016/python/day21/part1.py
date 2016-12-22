#!/usr/bin/python

with open('input.txt') as file_handler:
    instructions = [line.rstrip().split() for line in file_handler]

password = 'abcdefgh'

for instr in instructions:
    if instr[0] == 'swap':
        if instr[1] == 'position':
            index1 = int(instr[2])
            index2 = int(instr[5])
        elif instr[1] == 'letter':
            index1 = password.index(instr[2])
            index2 = password.index(instr[5])

        password = list(password)
        password[index1], password[index2] = password[index2], password[index1]
        password = ''.join(password)
    elif instr[0] == 'rotate':
        if instr[1] == 'based':
            index = password.index(instr[6])
            steps = -index - 1

            if index >= 4:
                steps -= 1
        elif instr[1] == 'right':
            steps = -int(instr[2])
        elif instr[1] == 'left':
            steps = int(instr[2])

        steps = steps % len(password)
        password = password[steps:] + password[:steps]
    elif instr[0] == 'reverse':
        index1 = int(instr[2])
        index2 = int(instr[4])

        password = password[:index1] + ''.join(
            reversed(password[index1: index2 + 1])
        ) + password[index2 + 1:]
    elif instr[0] == 'move':
        index1 = int(instr[2])
        index2 = int(instr[5])

        char = password[index1]
        password = (password[:index1] + password[index1 + 1:])
        password = password[:index2] + char + password[index2:]

print password