#!/usr/bin/ruby

code = 20151125
row = 2947
col = 3029

size = row + col - 1
iterations = size * (size + 1) / 2 - row

iterations.times do
    code = code * 252533 % 33554393
end

puts code
