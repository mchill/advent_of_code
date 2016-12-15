#!/usr/bin/python
import hashlib
import re


class Match:
    def __init__(self, index, search):
        self.index = index
        self.search = search
        self.found = False


matches = []
triple = re.compile(r'(\w)\1\1')

keys = 0
index = -1

while keys < 64:
    index += 1
    key = hashlib.md5('zpqevtbw' + str(index)).hexdigest()

    matches[:] = [match for match in matches if not match.found and index - match.index <= 1000]
    for match in matches:
        if match.search in key:
            match.found = True
            keys += 1

            if keys == 64:
                print match.index
                break

    match = re.search(triple, key)
    if match:
        matches.append(Match(index, match.group(1)[0] * 5))
