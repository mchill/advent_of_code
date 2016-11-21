#!/usr/bin/ruby
require 'json'

def iterate(value)
    if value.is_a?(Hash)
        return iterateHash(value)
    elsif value.is_a?(Array)
        return iterateArray(value)
    elsif value.is_a?(Integer)
        return value
    end

    return 0
end

def iterateHash(data)
    total = 0

    data.each do |key, value|
        if value == 'red'
            total = 0
            break
        end
        total += iterate(value)
    end

    return total
end

def iterateArray(data)
    total = 0

    data.each do |value|
        total += iterate(value)
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
