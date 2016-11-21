#!/usr/bin/ruby
require 'json'

def iterate(data)
    total = 0

    data.each do |value|
        if value.is_a?(Hash) or value.is_a?(Array)
            total += iterate(value)
        elsif value.is_a?(Integer)
            total += value
        end
    end

    return total
end

begin
    File.open('input.txt', 'r') do |file|
        data = JSON.parse(file.read)
        puts iterate(data)
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
