#!/usr/bin/ruby

def expand(variable, wires)
    if variable.to_i.to_s == variable
        return variable.to_i
    end
    equation = wires[variable]

    if equation.size == 1
        value = expand(equation[0], wires)
    elsif equation.size == 2
        value = ~expand(equation[1], wires)
    elsif equation[1] == 'AND'
        value = expand(equation[0], wires) & expand(equation[2], wires)
    elsif equation[1] == 'OR'
        value = expand(equation[0], wires) | expand(equation[2], wires)
    elsif equation[1] == 'LSHIFT'
        value = expand(equation[0], wires) << equation[2].to_i
    elsif equation[1] == 'RSHIFT'
        value = expand(equation[0], wires) >> equation[2].to_i
    end

    wires[variable] = [value.to_s]
    return value
end

begin
    File.open('input.txt', 'r') do |file|
        wires = {}

        file.each_line do |line|
            parts = line.strip.split(' -> ')
            wires[parts[1]] = parts[0].split(' ')
        end

        puts expand('a', wires)
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
