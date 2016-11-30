#!/usr/bin/ruby

minimum = 33100000
house = 0
gifts = 0

while gifts * 11 < minimum
    house += 1
    gifts = 0

    for elf in 1..Math.sqrt(house).to_i
        if house % elf == 0
            complement = house / elf

            if elf < 50
                gifts += complement
            end

            if complement < 50 and elf != complement
                gifts += elf
            end
        end
    end
end

puts house
