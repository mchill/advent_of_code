#!/usr/bin/python
from collections import deque
import hashlib


def valid_moves(current):
    x, y, path = current
    hash = hashlib.md5('bwnlcvfs' + path).hexdigest()
    moves = [(x, y - 1, path + 'U'), (x, y + 1, path + 'D'),
             (x - 1, y, path + 'L'), (x + 1, y, path + 'R')]
    return [move for ind, move in enumerate(moves) if valid(move, hash[ind])]


def valid(move, char):
    x, y, path = move
    return int(char, 16) > 10 and 0 <= x < 4 and 0 <= y < 4


moves = deque([(0, 0, '')])

while not any(x == 3 and y == 3 for x, y, _ in moves):
    moves.extend(valid_moves(moves.popleft()))

print next(path for x, y, path in moves if x == 3 and y == 3)
