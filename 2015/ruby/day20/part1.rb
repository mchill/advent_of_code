#!/usr/bin/ruby

minimum = 33100000 / 10
house = 0
gifts = 0

while gifts < minimum
    house += 1
    elves = (1..Math.sqrt(house).to_i).select { |n| house % n == 0 }
    gifts = elves.inject(0){ |sum, elf| sum + (elf + house / elf) }
end

puts house
