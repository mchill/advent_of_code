#!/usr/bin/python
import hashlib

door = 'ffykfhsq'
index = 0
password = [None] * 8

while None in password:
    digest = hashlib.md5(door + str(index)).hexdigest()
    index += 1

    if digest.startswith('00000'):
        position = int(digest[5], 16)

        if position < 8 and password[position] is None:
            password[position] = digest[6]

print ''.join(password)
