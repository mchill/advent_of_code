#!/usr/bin/python
import sys

with open('input.txt') as file_handler:
    lines = file_handler.readlines()

bots = {}
outputs = {}

for line in lines:
    line = line.rstrip().split()

    if line[0] == 'value':
        if line[5] not in bots:
            bots[line[5]] = { 'values': [] }
        bots[line[5]]['values'].append(int(line[1]))
        continue

    if line[1] not in bots:
        bots[line[1]] = { 'values': [] }

    bots[line[1]]['low'] = [line[5], line[6]]
    bots[line[1]]['high'] = [line[10], line[11]]

while bots:
    complete = { key: bot for (key, bot) in bots.items() if len(bot['values']) == 2 }
    bots = { key: bot for (key, bot) in bots.items() if len(bot['values']) < 2 }

    for key, bot in complete.items():
        if bot['low'][0] == 'bot':
            bots[bot['low'][1]]['values'].append(min(bot['values']))
        else:
            outputs[bot['low'][1]] = min(bot['values'])

        if bot['high'][0] == 'bot':
            bots[bot['high'][1]]['values'].append(max(bot['values']))
        else:
            outputs[bot['high'][1]] = max(bot['values'])

print outputs['0'] * outputs['1'] * outputs['2']
