#!/usr/bin/python

with open('input.txt') as file_handler:
    ranges = [map(int, line.rstrip().split('-')) for line in file_handler]

ip = 0
allowed = 0

while ip <= 4294967295:
    upper = next((rng[1] for rng in ranges if rng[0] <= ip <= rng[1]), None)

    if upper:
        ip = upper + 1
        continue

    allowed += 1
    ip += 1

print allowed
