#!/usr/bin/python


def parse(value, registers):
    return registers[value] if value in registers else int(value)


with open('input.txt') as file_handler:
    instructions = [line.rstrip().split() for line in file_handler]

a = 0
out = ''

while out != '01010101':
    registers = {'a': a, 'b': 0, 'c': 0, 'd': 0}
    index = 0

    a += 1
    out = ''

    while index < len(instructions):
        instr = instructions[index]
        index += 1

        if instr[0] == 'cpy':
            registers[instr[2]] = parse(instr[1], registers)
        elif instr[0] == 'inc':
            registers[instr[1]] += 1
        elif instr[0] == 'dec':
            registers[instr[1]] -= 1
        elif instr[0] == 'out':
            out += str(parse(instr[1], registers))
            if len(out) == 8:
                break
        elif parse(instr[1], registers) != 0:
            index += parse(instr[2], registers) - 1

print a - 1
