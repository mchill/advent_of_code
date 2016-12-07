#!/usr/bin python

with open('input.txt') as fh:
    print ''.join([min(set(ch), key=ch.count) for ch in zip(*fh.readlines())])
