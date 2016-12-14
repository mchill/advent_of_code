#!/usr/bin/python


def valid_moves(current, traversed):
    x, y = current

    moves = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]
    return [move for move in moves if valid(move, traversed)]


def valid(position, traversed):
    x, y = position

    if x < 0 or y < 0 or position in traversed:
        return False

    if bin(x*x + 3*x + 2*x*y + y + y*y + 1364).count('1') % 2 == 0:
        traversed.add(position)
        return True


moves = [(1, 1)]
traversed = {(1, 1)}

for _ in range(50):
    old_moves = moves
    moves = []

    for move in old_moves:
        moves.extend(valid_moves(move, traversed))

print len(traversed)
