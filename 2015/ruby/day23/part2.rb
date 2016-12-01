#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        instructions = []
        registers = {'a' => 1, 'b' => 0}
        index = 0

        file.each_line do |line|
            instructions.push(line.tr(',', '').chomp)
        end

        while index < instructions.count
            instruction = instructions[index].split

            case instruction[0]
            when 'hlf'
                registers[instruction[1]] /= 2
            when 'tpl'
                registers[instruction[1]] *= 3
            when 'inc'
                registers[instruction[1]] += 1
            when 'jmp'
                index += instruction[1].to_i
                next
            when 'jie'
                if registers[instruction[1]].even?
                    index += instruction[2].to_i - 1
                end
            when 'jio'
                if registers[instruction[1]] == 1
                    index += instruction[2].to_i - 1
                end
            end

            index += 1
        end

        puts registers['b']
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
