#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        nice = 0

        file.each_line do |line|
            if not (line.match(/(\w\w).*\1/).nil? or
                    line.match(/(\w)\w\1/).nil?)
                nice += 1
            end
        end

        puts nice
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
