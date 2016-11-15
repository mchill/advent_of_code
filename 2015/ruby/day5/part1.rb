#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        nice = 0

        file.each_line do |line|
            if not (line.scan(/[aeiou]/).count < 3 or
                    line.scan(/(\w)\1+/).count < 1 or
                    line.include?('ab') or
                    line.include?('cd') or
                    line.include?('pq') or
                    line.include?('xy'))
                nice += 1
            end
        end

        puts nice
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
