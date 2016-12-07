#!/usr/bin python
import re


def is_supported(address):
    address = re.split(r'\[|\]', address)

    supernet = ','.join(address[::2])
    hypernet = ','.join(address[1::2])

    matches = re.findall(r'(\w)(?=(?!\1)(\w\1))', supernet)
    abas = [''.join(match) for match in matches]

    return True if any(aba[1:] + aba[1] in hypernet for aba in abas) else False


with open('input.txt') as file_handler:
    addresses = file_handler.readlines()

print sum(1 for address in addresses if is_supported(address))
