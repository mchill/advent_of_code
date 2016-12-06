#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.readlines()

sector_id_sum = 0

for line in content:
    data = line.rstrip(']\n').rsplit('-', 1)
    name = data[0].translate(None, '-')
    sector_id, checksum = data[1].split('[')

    most_common = sorted((-name.count(letter), letter) for letter in set(name))
    if checksum == ''.join(letter for _, letter in most_common[:5]):
        sector_id_sum += int(sector_id)

print sector_id_sum
