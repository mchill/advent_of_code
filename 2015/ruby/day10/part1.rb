#!/usr/bin/ruby

sequence = '3113322113'

for i in 1..40
    new = ''
    currentChar = sequence[0]
    count = 0

    sequence.split('').each do |char|
        if char != currentChar
            new << count.to_s << currentChar
            currentChar = char
            count = 0
        end

        count += 1
    end

    new << count.to_s << currentChar
    sequence = new
end

puts sequence.length
