#!/usr/bin python

with open('input.txt') as fh:
    print ''.join([max(set(ch), key=ch.count) for ch in zip(*fh.readlines())])
