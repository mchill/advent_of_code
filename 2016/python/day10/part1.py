#!/usr/bin/python
import sys

with open('input.txt') as file_handler:
    lines = file_handler.readlines()

bots = {}

for line in lines:
    line = line.rstrip().split()

    if line[0] == 'value':
        if line[5] not in bots:
            bots[line[5]] = { 'values': [] }
        bots[line[5]]['values'].append(int(line[1]))
        continue

    if line[1] not in bots:
        bots[line[1]] = { 'values': [] }

    if line[5] == 'bot':
        bots[line[1]]['low'] = line[6]
    if line[10] == 'bot':
        bots[line[1]]['high'] = line[11]

while bots:
    complete = { key: bot for (key, bot) in bots.items() if len(bot['values']) == 2 }
    bots = { key: bot for (key, bot) in bots.items() if len(bot['values']) < 2 }

    for key, bot in complete.items():
        if 61 in bot['values'] and 17 in bot['values']:
            print key
            sys.exit(0)

        if 'low' in bot:
            bots[bot['low']]['values'].append(min(bot['values']))
        if 'high' in bot:
            bots[bot['high']]['values'].append(max(bot['values']))
