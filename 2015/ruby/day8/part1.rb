#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        total = 0

        file.each_line do |line|
            string = line[1..-3]
            total += 2

            newString = string.gsub('\\"', '').gsub('\\\\', '')
            total += (string.length - newString.length) / 2
            total += newString.scan(/\\x/).count * 3
        end

        puts total
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
