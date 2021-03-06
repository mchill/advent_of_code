#!/usr/bin/ruby

def compare(type, goalVal, testVal)
    if type == 'cats' or type == 'trees'
        if testVal > goalVal
            return true
        end
        return false
    elsif type == 'pomeranians' or type == 'goldfish'
        if testVal < goalVal
            return true
        end
        return false
    elsif testVal == goalVal
        return true
    end
    return false
end

begin
    File.open('input.txt', 'r') do |file|
        goal = {
            'children' => 3,
            'cats' => 7,
            'samoyeds' => 2,
            'pomeranians' => 3,
            'akitas' => 0,
            'vizslas' => 0,
            'goldfish' => 5,
            'trees' => 3,
            'cars' => 2,
            'perfumes' => 1
        }

        file.each_line do |line|
            stats = line.delete(':,').downcase.split(' ')

            if (compare(stats[2], goal[stats[2]], stats[3].to_i) and
                compare(stats[4], goal[stats[4]], stats[5].to_i) and
                compare(stats[6], goal[stats[6]], stats[7].to_i))
                puts stats[1]
                break
            end
        end
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
