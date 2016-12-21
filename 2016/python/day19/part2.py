#!/usr/bin/python

INITIAL = 3004953


class Elf:
    def __init__(self, id):
        self.id = id


across = Elf(1)
elf = across

for id in range(2, INITIAL + 1):
    elf.next = Elf(id)
    elf = elf.next

elf.next = across

for _ in range(INITIAL / 2 - 1):
    across = across.next

for round in range(INITIAL):
    across.next = across.next.next

    if round & 1 == 0:
        across = across.next

print across.id
