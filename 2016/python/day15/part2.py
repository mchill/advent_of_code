#!/usr/bin/python


def parse(disc):
    return {'position': int(disc[11]), 'size': int(disc[3])}


with open('input.txt') as file_handler:
    discs = [parse(line.rstrip('.\n').split()) for line in file_handler]

discs.append({'position': 0, 'size': 11})
time = 0

while True:
    for index, disc in enumerate(discs):
        if (disc['position'] + index + time + 1) % disc['size'] != 0:
            break
    else:
        print time
        break

    time += 1
