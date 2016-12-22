#!/usr/bin/python

with open('input.txt') as file_handler:
    ranges = [map(int, line.rstrip().split('-')) for line in file_handler]

ip = 0
upper = 0

while upper is not None:
    ip = upper + 1
    upper = next((rng[1] for rng in ranges if rng[0] <= ip <= rng[1]), None)

print ip
