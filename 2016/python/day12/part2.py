#!/usr/bin/python


def parse(value, registers):
    return registers[value] if value in registers else int(value)


with open('input.txt') as file_handler:
    instructions = [line.rstrip().split() for line in file_handler]

registers = {'a': 0, 'b': 0, 'c': 1, 'd': 0}
index = 0

while index < len(instructions):
    instr = instructions[index]
    index += 1

    if instr[0] == 'cpy':
        registers[instr[2]] = parse(instr[1], registers)
    elif instr[0] == 'inc':
        registers[instr[1]] += 1
    elif instr[0] == 'dec':
        registers[instr[1]] -= 1
    elif parse(instr[1], registers) != 0:
        index += int(instr[2]) - 1

print registers['a']
