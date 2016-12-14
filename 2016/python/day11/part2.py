#!/usr/bin/python
from itertools import combinations
from collections import deque


PAIRS = 7


class Generator:
    def __init__(self, element):
        self.element = element


class Microchip:
    def __init__(self, element, match):
        self.element = element
        self.match = match


class State:
    def __init__(self, floor, floors, csf):
        self.floor = floor
        self.floors = floors
        self.csf = csf
        self.key = self.key()

    def key(self):
        key = []

        for element in range(PAIRS):
            key.append((
                next(index for index, floor in enumerate(self.floors) if any(
                    isinstance(part, Generator) and part.element == element for part in floor
                )),
                next(index for index, floor in enumerate(self.floors) if any(
                    isinstance(part, Microchip) and part.element == element for part in floor
                ))
            ))

        return (self.floor, tuple(sorted(key)))

    def moves(self, traversed):
        combos = (list(combinations(self.floors[self.floor], 2)) +
                  list(combinations(self.floors[self.floor], 1)))
        valid_moves = []

        for direction in [1, -1]:
            next_floor = self.floor + direction
            if not 0 <= next_floor < 4:
                continue

            for combo in combos:
                floors = [list(floor) for floor in self.floors]

                for part in combo:
                    floors[next_floor].append(part)
                    floors[self.floor].remove(part)

                valid_moves.append(State(next_floor, floors, self.csf + 1))

        return [move for move in valid_moves if move.valid(traversed)]

    def valid(self, traversed):
        if self.key in traversed:
            traversed[self.key].csf = min(traversed[self.key].csf, self.csf)
            return False

        for floor in self.floors:
            if (any(isinstance(part, Microchip) and part.match not in floor for part in floor) and
                any(isinstance(part, Generator) for part in floor)):
                return False

        traversed[self.key] = self
        return True


def main():
    gen = []
    chip = []

    for pair in range(PAIRS):
        gen.append(Generator(pair))
        chip.append(Microchip(pair, gen[pair]))

    initial = State(0, [
        [gen[0], chip[0], gen[5], chip[5], gen[6], chip[6]],
        [gen[1], gen[2], gen[3], gen[4]],
        [chip[1], chip[2], chip[3], chip[4]],
        []
    ], 0)

    traversed = {}
    possible = deque([initial])

    while len(possible):
        current = possible.popleft()

        if len(current.floors[0]) + len(current.floors[1]) + len(current.floors[2]) == 0:
            print current.csf
            break

        possible.extend(current.moves(traversed))


if __name__ == "__main__":
    main()
