#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        totalLength = 0

        file.each_line do |line|
            dimensions = line.split('x')

            width = dimensions[0].to_i
            height = dimensions[1].to_i
            length = dimensions[2].to_i

            volume = width * height * length
            perimeter1 = (width + height) * 2
            perimeter2 = (width + length) * 2
            perimeter3 = (height + length) * 2

            totalLength += volume + [perimeter1, perimeter2, perimeter3].min
        end

        puts totalLength
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
