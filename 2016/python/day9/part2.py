#!/usr/bin/python


def decompress(string):
    length = 0

    while '(' in string:
        parts = string.split('(', 1)
        marker, string = parts[1].split(')', 1)
        size, repetitions = map(int, marker.split('x'))

        length += len(parts[0]) + decompress(string[:size]) * repetitions
        string = string[size:]

    length += len(string)
    return length


with open('input.txt') as file_handler:
    print decompress(file_handler.read().rstrip())
