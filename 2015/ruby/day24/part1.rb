#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        packages = []

        file.each_line do |line|
            packages.push(line.to_i)
        end

        third = packages.inject(:+) / 3
        min_entanglement = nil

        for first_size in 1..packages.count
            first_groups = packages.combination(first_size)

            first_groups.each do |first_group|
                if first_group.inject(:+) != third then next end

                entanglement = first_group.inject(:*)

                if min_entanglement.nil? or entanglement < min_entanglement
                    min_entanglement = entanglement
                end
            end

            if not min_entanglement.nil? then break end
        end

        puts min_entanglement
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
