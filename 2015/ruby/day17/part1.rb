#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        containers = []
        combinations = []
        fit = 0

        file.each_line do |line|
            containers.push(line.to_i)
        end

        for i in 1..containers.count
            combinations.concat(containers.combination(i).to_a)
        end

        combinations.each do |combination|
            if combination.inject(:+) == 150
                fit += 1
            end
        end

        puts fit
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
