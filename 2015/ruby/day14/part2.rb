#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        reindeer = []

        file.each_line do |line|
            sentence = line.split(' ')
            reindeer.push({
                'speed' => sentence[3].to_i,
                'flyTime' => sentence[6].to_i,
                'restTime' => sentence[13].to_i,
                'distance' => 0,
                'flying' => true,
                'timeLeft' => sentence[6].to_i,
                'score' => 0
            })
        end

        for i in 1..2503
            reindeer.each do |stats|
                stats['timeLeft'] -= 1

                if stats['flying']
                    stats['distance'] += stats['speed']

                    if stats['timeLeft'] == 0
                        stats['flying'] = false
                        stats['timeLeft'] = stats['restTime']
                    end

                    next
                end

                if stats['timeLeft'] == 0
                    stats['flying'] = true
                    stats['timeLeft'] = stats['flyTime']
                end
            end

            reindeer.group_by{ |value| value['distance'] }.max.last.each do |stats|
                stats['score'] += 1
            end
        end

        puts reindeer.max_by{ |value| value['score'] }['score']
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
