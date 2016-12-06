#!/usr/bin/python

with open('input.txt') as file_handler:
    content = file_handler.readlines()

for line in content:
    data = line.rstrip(']\n').rsplit('-', 1)
    name = data[0].translate(None, '-')
    sector_id = data[1].split('[')[0]
    name = ''.join(chr((ord(letter) - 97 + int(sector_id)) % 26 + 97) for letter in name)

    if name == 'northpoleobjectstorage':
        print sector_id
        break
