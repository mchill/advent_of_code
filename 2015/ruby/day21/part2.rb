#!/usr/bin/ruby

boss = {
    :health => 104,
    :damage => 8,
    :armor => 1
}

weapons = [
    {:cost => 8,  :damage => 4},
    {:cost => 10, :damage => 5},
    {:cost => 25, :damage => 6},
    {:cost => 40, :damage => 7},
    {:cost => 74, :damage => 8}
]

armor = [
    {:cost => 0,   :armor => 0},
    {:cost => 13,  :armor => 1},
    {:cost => 31,  :armor => 2},
    {:cost => 53,  :armor => 3},
    {:cost => 75,  :armor => 4},
    {:cost => 102, :armor => 5}
]

rings = [
    {:cost => 0,   :damage => 0, :armor => 0},
    {:cost => 0,   :damage => 0, :armor => 0},
    {:cost => 25,  :damage => 1, :armor => 0},
    {:cost => 50,  :damage => 2, :armor => 0},
    {:cost => 100, :damage => 3, :armor => 0},
    {:cost => 20,  :damage => 0, :armor => 1},
    {:cost => 40,  :damage => 0, :armor => 2},
    {:cost => 80,  :damage => 0, :armor => 3}
]

most_spent = 0

weapons.each do |weapon|
    armor.each do |armor|
        rings.each do |ring1|
            rings.each do |ring2|
                if ring1[:cost] == ring2[:cost] and ring1[:cost] != 0
                    next
                end

                player = {
                    :health => 100,
                    :damage => weapon[:damage] + ring1[:damage] + ring2[:damage],
                    :armor => armor[:armor] + ring1[:armor] + ring2[:armor]
                }

                hits_to_kill_player = (player[:health] - 1) / [boss[:damage] - player[:armor], 1].max
                hits_to_kill_boss = (boss[:health] - 1) / [player[:damage] - boss[:armor], 1].max

                if hits_to_kill_player < hits_to_kill_boss
                    cost = weapon[:cost] + armor[:cost] + ring1[:cost] + ring2[:cost]
                    most_spent = cost > most_spent ? cost : most_spent
                end
            end
        end
    end
end

puts most_spent
