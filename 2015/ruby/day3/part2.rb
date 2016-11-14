#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        position1 = [0, 0]
        position2 = [0, 0]
        visited = [[0, 0]]

        file.each_char.with_index do |char, index|
            position = position1
            if index.odd? then
                position = position2
            end

            case char
            when '^'
                position[1] -= 1
            when '>'
                position[0] += 1
            when 'v'
                position[1] += 1
            when '<'
                position[0] -= 1
            end

            if (visited.detect {|x, y| x == position[0] and y == position[1]}).nil? then
                visited.push(position.clone)
            end
        end

        puts visited.count
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
