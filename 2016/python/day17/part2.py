#!/usr/bin/python
from collections import deque
import hashlib


def valid_moves(current, complete):
    x, y, path = current
    hash = hashlib.md5('bwnlcvfs' + path).hexdigest()
    moves = [(x, y - 1, path + 'U'), (x, y + 1, path + 'D'),
             (x - 1, y, path + 'L'), (x + 1, y, path + 'R')]
    return [move for ind, move in enumerate(moves)
            if valid(move, hash[ind], complete)]


def valid(move, char, complete):
    x, y, path = move

    if int(char, 16) <= 10:
        return False

    if x == 3 and y == 3:
        complete.append(path)
        return False

    return 0 <= x < 4 and 0 <= y < 4


moves = deque([(0, 0, '')])
complete = []

while len(moves) > 0:
    moves.extend(valid_moves(moves.popleft(), complete))

print len(complete[-1])
