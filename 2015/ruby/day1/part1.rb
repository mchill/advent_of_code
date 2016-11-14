#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        floor = 0

        file.each_char do |char|
            case char
            when '('
                floor += 1
            when ')'
                floor -= 1
            end
        end

        puts floor
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
