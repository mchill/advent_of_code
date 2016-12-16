#!/usr/bin/python

data = '10010000000110000'
length = 272

while len(data) < length:
    data += '0' + ''.join('1' if char == '0' else '0' for char in data[::-1])

checksum = data[:length]

while len(checksum) % 2 == 0:
    checksum = ['1' if checksum[index] == checksum[index + 1]
                else '0' for index in range(0, len(checksum), 2)]

print ''.join(checksum)
