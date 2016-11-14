#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        floor = 0

        file.each_char.with_index do |char, index|
            case char
            when '('
                floor += 1
            when ')'
                floor -= 1
            end

            if floor < 0 then
                puts index + 1
                break
            end
        end
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
