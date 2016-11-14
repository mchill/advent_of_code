#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        area = 0

        file.each_line do |line|
            dimensions = line.split('x')

            width = dimensions[0].to_i
            height = dimensions[1].to_i
            length = dimensions[2].to_i

            side1 = width * height
            side2 = width * length
            side3 = height * length

            area += 2 * (side1 + side2 + side3) + [side1, side2, side3].min
        end

        puts area
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
