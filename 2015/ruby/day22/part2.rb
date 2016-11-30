#!/usr/bin/ruby

def effects(boss, player)
    if player[:shield] == 0
        player[:armor] = 0
    end

    if player[:poison] > 0
        boss[:health] -= 3
    end

    if player[:recharge] > 0
        player[:mana] += 101
    end

    player[:shield] = [player[:shield] - 1, 0].max
    player[:poison] = [player[:poison] - 1, 0].max
    player[:recharge] = [player[:recharge] - 1, 0].max
end

def turn(old_boss, old_player)
    least_used = 9001

    for move in 1..5
        boss = old_boss.clone
        player = old_player.clone

        player[:health] -= 1
        if player[:health] <= 0 then next end

        # Effects
        effects(boss, player)
        if boss[:health] <= 0
            least_used = [player[:used_mana], least_used].min
            next
        end

        # Player attacks
        case move
        when 1
            # Magic Missile
            player[:mana] -= 53
            player[:used_mana] += 53
            boss[:health] -= 4
        when 2
            # Drain
            player[:mana] -= 73
            player[:used_mana] += 73
            boss[:health] -=2
            player[:health] += 2
        when 3
            # Shield
            if player[:shield] > 0 then next end
            player[:mana] -= 113
            player[:used_mana] += 113
            player[:shield] = 6
            player[:armor] = 7
        when 4
            # Poison
            if player[:poison] > 0 then next end
            player[:mana] -= 173
            player[:used_mana] += 173
            player[:poison] = 6
        when 5
            # Recharge
            if player[:recharge] > 0 then next end
            player[:mana] -= 229
            player[:used_mana] += 229
            player[:recharge] = 5
        end

        if player[:mana] < 0 then next end

        # Effects
        effects(boss, player)

        # If boss died
        if boss[:health] <= 0
            least_used = [player[:used_mana], least_used].min
            next
        end

        # Boss attacks
        player[:health] -= [boss[:damage] - player[:armor], 1].max

        # If player died
        if player[:health] <= 0 then next end

        least_used = [turn(boss, player), least_used].min
    end

    return least_used
end

boss = {
    :health => 71,
    :damage => 10
}

player = {
    :health => 50,
    :mana => 500,
    :used_mana => 0,
    :armor => 0,
    :shield => 0,
    :poison => 0,
    :recharge => 0
}

puts turn(boss, player)
