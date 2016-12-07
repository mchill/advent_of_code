#!/usr/bin python
import re

with open('input.txt') as file_handler:
    addrs = file_handler.readlines()

good = re.compile(r'(\w)(?!\1)(\w)\2\1')
bad = re.compile(r'\[\w*(\w)(?!\1)(\w)\2\1\w*\]')

print sum(1 for addr in addrs if good.search(addr) and not bad.search(addr))
