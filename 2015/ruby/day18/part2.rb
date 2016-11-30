#!/usr/bin/ruby

SIZE = 100

begin
    File.open('input.txt', 'r') do |file|
        lights = [Array.new(SIZE + 2){ |i| 0 }]

        file.each_line do |line|
            row = [0]

            line.split('').each do |char|
                if char == '#'
                    row.push(1)
                    next
                end
                row.push(0)
            end

            lights.push(row)
        end

        lights[1][1] = 1
        lights[1][SIZE] = 1
        lights[SIZE][1] = 1
        lights[SIZE][SIZE] = 1

        lights.push(Array.new(SIZE + 2){ |i| 0 })
        new_lights = lights

        for frame in 1..100
            lights = Marshal.load(Marshal.dump(new_lights))

            for row in 1..SIZE
                for col in 1..SIZE
                    surrounding = (lights[row][col - 1] + lights[row - 1][col - 1] +
                                   lights[row - 1][col] + lights[row - 1][col + 1] +
                                   lights[row][col + 1] + lights[row + 1][col + 1] +
                                   lights[row + 1][col] + lights[row + 1][col - 1])

                    if (row == 1 or row == SIZE) and (col == 1 or col == SIZE)
                        new_lights[row][col] = 1
                    elsif lights[row][col] == 1 and surrounding != 2 and surrounding != 3
                        new_lights[row][col] = 0
                    elsif lights[row][col] == 0 and surrounding == 3
                        new_lights[row][col] = 1
                    end
                end
            end
        end

        puts new_lights.flatten.count(1)
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
