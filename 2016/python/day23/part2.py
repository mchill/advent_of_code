#!/usr/bin/python


def parse(value, registers):
    return registers[value] if value in registers else int(value)


with open('input.txt') as file_handler:
    instructions = [line.rstrip().split() for line in file_handler]

swaps = {'cpy': 'jnz', 'jnz': 'cpy', 'inc': 'dec', 'dec': 'inc', 'tgl': 'inc'}
registers = {'a': 12, 'b': 0, 'c': 0, 'd': 0}
index = 0

while index < len(instructions):
    instr = instructions[index]
    index += 1

    if index == 4:
        registers['a'] = registers['b'] * registers['d']
        registers['c'] = 0
        registers['d'] = 0
        index = 10
    elif instr[0] == 'cpy' and isinstance(instr[1], basestring):
        registers[instr[2]] = parse(instr[1], registers)
    elif instr[0] == 'inc' and isinstance(instr[1], basestring):
        registers[instr[1]] += 1
    elif instr[0] == 'dec' and isinstance(instr[1], basestring):
        registers[instr[1]] -= 1
    elif instr[0] == 'jnz' and parse(instr[1], registers) != 0:
        index += parse(instr[2], registers) - 1
    elif instr[0] == 'tgl':
        try:
            toggle = instructions[index - 1 + parse(instr[1], registers)]
            toggle[0] = swaps[toggle[0]]
        except IndexError:
            continue

print registers['a']
