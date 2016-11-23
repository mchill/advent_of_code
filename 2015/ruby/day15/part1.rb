#!/usr/bin/ruby

begin
    File.open('input.txt', 'r') do |file|
        ingredients = {}
        max = 0

        file.each_line do |line|
            stats = line.delete(':,').downcase.split(' ')

            ingredients[stats[0]] = {
                :capacity => stats[2].to_i,
                :durability => stats[4].to_i,
                :flavor => stats[6].to_i,
                :texture => stats[8].to_i
            }
        end

        for sugar in 0..100
            for sprinkles in 0..(100 - sugar)
                for candy in 0..(100 - sugar - sprinkles)
                    chocolate = 100 - sugar - sprinkles - candy

                    capacity = (
                        sugar * ingredients['sugar'][:capacity] +
                        sprinkles * ingredients['sprinkles'][:capacity] +
                        candy * ingredients['candy'][:capacity] +
                        chocolate * ingredients['chocolate'][:capacity]
                    )
                    durability = (
                        sugar * ingredients['sugar'][:durability] +
                        sprinkles * ingredients['sprinkles'][:durability] +
                        candy * ingredients['candy'][:durability] +
                        chocolate * ingredients['chocolate'][:durability]
                    )
                    flavor = (
                        sugar * ingredients['sugar'][:flavor] +
                        sprinkles * ingredients['sprinkles'][:flavor] +
                        candy * ingredients['candy'][:flavor] +
                        chocolate * ingredients['chocolate'][:flavor]
                    )
                    texture = (
                        sugar * ingredients['sugar'][:texture] +
                        sprinkles * ingredients['sprinkles'][:texture] +
                        candy * ingredients['candy'][:texture] +
                        chocolate * ingredients['chocolate'][:texture]
                    )

                    capacity = capacity < 0 ? 0 : capacity
                    durability = durability < 0 ? 0 : durability
                    flavor = flavor < 0 ? 0 : flavor
                    texture = texture < 0 ? 0 : texture

                    total = capacity * durability * flavor * texture
                    max = total > max ? total : max
                end
            end
        end

        puts max
    end
rescue SystemCallError
    puts 'input.txt not found.'
end
