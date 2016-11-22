#!/usr/bin/ruby

def sit(sat, happiness)
    max = 0

    for i in 0..8
        next if sat.include?(i)
        max = [sit(sat.push(i), happiness), max].max
        sat.pop()
    end

    if sat.size == 9
        max += happiness[sat[0]][sat[8]] + happiness[sat[8]][sat[0]]
        for i in 1..8
            max += happiness[sat[i]][sat[i - 1]] + happiness[sat[i - 1]][sat[i]]
        end
    end

    return max
end

begin
    File.open('input.txt', 'r') do |file|
        people = {
            'Alice' => 0,
            'Bob' => 1,
            'Carol' => 2,
            'David' => 3,
            'Eric' => 4,
            'Frank' => 5,
            'George' => 6,
            'Mallory' => 7
        }
        happiness = Array.new(9) { |value| Array.new(9) { |value| 0 }}

        file.each_line do |line|
            sentence = line[0..-3].split(' ')
            happiness[people[sentence[0]]][people[sentence[10]]] = sentence[3].to_i

            if sentence[2] == 'lose'
                happiness[people[sentence[0]]][people[sentence[10]]] *= -1
            end
        end

        puts sit([], happiness)
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
