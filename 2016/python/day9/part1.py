#!/usr/bin/python

with open('input.txt') as file_handler:
    compressed = file_handler.read().rstrip()

decompressed = ''

while compressed:
    if compressed[0] != '(':
        decompressed += compressed[0]
        compressed = compressed[1:]
        continue

    marker, compressed = compressed.split(')', 1)
    size, repetitions = map(int, marker[1:].split('x'))

    decompressed += compressed[:size] * repetitions
    compressed = compressed[size:]

print len(decompressed)
