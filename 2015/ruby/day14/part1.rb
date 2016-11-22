#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        reindeer = []
        max = 0

        file.each_line do |line|
            sentence = line.split(' ')
            reindeer.push({
                'speed' => sentence[3].to_i,
                'flyTime' => sentence[6].to_i,
                'restTime' => sentence[13].to_i
            })
        end

        reindeer.each do |stats|
            period = stats['flyTime'] + stats['restTime']
            flightTime = (2503 / period) * stats['flyTime']
            flightTime += [2503 % period, stats['flyTime']].min
            max = [flightTime * stats['speed'], max].max
        end

        puts max
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
