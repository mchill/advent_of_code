#!/usr/bin/python
import hashlib

door = 'ffykfhsq'
index = 0
password = ''

while len(password) < 8:
    digest = hashlib.md5(door + str(index)).hexdigest()
    index += 1

    if digest.startswith('00000'):
        password += digest[5]

print password
