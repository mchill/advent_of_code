#!/usr/bin/ruby

def travel(path, distances)
    shortest = nil

    for i in 0..7
        next if path.include?(i)

        distance = travel(path.push(i), distances)
        path.pop()

        if path.size > 0
            distance += distances[path[path.size - 1]][i]
        end

        shortest = [shortest, distance].compact.min
    end

    return shortest || 0
end

begin
    File.open('input.txt', 'r') do |file|
        locations = {
            'AlphaCentauri' => 0,
            'Snowdin' => 1,
            'Tambi' => 2,
            'Faerun' => 3,
            'Norrath' => 4,
            'Straylight' => 5,
            'Tristram' => 6,
            'Arbre' => 7
        }
        distances = Array.new(8) { Array.new(8) }

        file.each_line do |line|
            line = line.split(' ')

            source = line[0]
            destination = line[2]
            distance = line[4].to_i

            distances[locations[source]][locations[destination]] = distance
            distances[locations[destination]][locations[source]] = distance
        end

        puts travel([], distances)
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
