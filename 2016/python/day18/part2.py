#!/usr/bin/python

with open('input.txt') as file_handler:
    row = '.' + file_handler.read().rstrip() + '.'

safe = row.count('.') - 2
traps = ['^^.', '.^^', '^..', '..^']

for _ in range(399999):
    next_row = '.'

    for index in range(1, len(row) - 1):
        next_row += '^' if row[index - 1: index + 2] in traps else '.'

    row = next_row + '.'
    safe += row.count('.') - 2

print safe
