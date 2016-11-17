#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        total = 0

        file.each_line do |line|
            string = line[1..-3]
            total += 4
            total += string.count('"\\')
        end

        puts total
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
