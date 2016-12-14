#!/usr/bin/python
from collections import deque


def valid_moves(current, traversed):
    x, y = current
    csf = traversed[current]

    moves = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]
    return [move for move in moves if valid(move, csf + 1, traversed)]


def valid(position, csf, traversed):
    x, y = position

    if x < 0 or y < 0:
        return False

    if position in traversed:
        traversed[position] = min(traversed[position], csf)
        return False

    if bin(x*x + 3*x + 2*x*y + y + y*y + 1364).count('1') % 2 == 0:
        traversed[position] = csf
        return True


moves = deque([(1, 1)])
traversed = {(1, 1): 0}

while (31, 39) not in traversed:
    moves.extend(valid_moves(moves.popleft(), traversed))

print traversed[(31, 39)]
