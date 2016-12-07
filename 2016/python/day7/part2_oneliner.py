#!/usr/bin python
import re

with open('input.txt') as fh:
    print sum(1 for addr in fh.readlines() if any(aba[1:] + aba[1] in ','.join(re.split(r'\[|\]', addr)[1::2]) for aba in [''.join(match) for match in re.findall(r'(\w)(?=(?!\1)(\w\1))', ','.join(re.split(r'\[|\]', addr)[::2]))]))
