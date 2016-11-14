#include <iostream>
#include <vector>
#include <stdlib.h>

using namespace std;

struct playerStruct {
    int health;
    int mana;
    int usedMana;
    int shield;
    int poison;
    int recharge;
};

struct bossStruct {
    int health;
    int damage;
};

int turn(playerStruct &player, bossStruct &boss) {
    int record = 9000;
    
    player.health--;
    if (player.health <= 0)
        return record;
    
    // Apply effects
    if (player.poison > 0)
        boss.health -= 3;
    if (player.recharge > 0)
        player.mana += 101;
    
    // Decrement timers
    player.shield = max(player.shield - 1, 0);
    player.poison = max(player.poison - 1, 0);
    player.recharge = max(player.recharge - 1, 0);
    
    for (int move = 0; move < 5; move++) {
        playerStruct nextPlayer = player;
        bossStruct nextBoss = boss;
        
        // Player casts a spell
        if (move == 0) {
            if (nextPlayer.mana < 53)
                break;
            
            nextPlayer.mana -= 53;
            nextPlayer.usedMana += 53;
            nextBoss.health -= 4;
        } else if (move == 1) {
            if (nextPlayer.mana < 73)
                break;
            
            nextPlayer.mana -= 73;
            nextPlayer.usedMana += 73;
            nextBoss.health -= 2;
            nextPlayer.health += 2;
        } else if (move == 2) {
            if (nextPlayer.mana < 113)
                break;
            if (nextPlayer.shield > 0)
                continue;
            
            nextPlayer.mana -= 113;
            nextPlayer.usedMana += 113;
            nextPlayer.shield = 6;
        } else if (move == 3) {
            if (nextPlayer.mana < 173)
                break;
            if (nextPlayer.poison > 0)
                continue;
            
            nextPlayer.mana -= 173;
            nextPlayer.usedMana += 173;
            nextPlayer.poison = 6;
        } else if (move == 4) {
            if (nextPlayer.mana < 229)
                break;
            if (nextPlayer.recharge > 0)
                continue;
            
            nextPlayer.mana -= 229;
            nextPlayer.usedMana += 229;
            nextPlayer.recharge = 5;
        }
        
        // Apply effects
        if (nextPlayer.poison > 0)
            nextBoss.health -= 3;
        if (nextPlayer.recharge > 0)
            nextPlayer.mana += 101;
        
        // Decrement timers
        nextPlayer.shield = max(nextPlayer.shield - 1, 0);
        nextPlayer.poison = max(nextPlayer.poison - 1, 0);
        nextPlayer.recharge = max(nextPlayer.recharge - 1, 0);
        
        if (nextBoss.health <= 0) {
            record = min(record, nextPlayer.usedMana);
            continue;
        }
    
        // Boss turn
        nextPlayer.health -= (10 - (nextPlayer.shield > 0) * 7);
        if (nextPlayer.health <= 0)
            continue;
    
        // Next turn
        record = min(record, turn(nextPlayer, nextBoss));
    }
    
    return record;
}

int main () {
    playerStruct player;
    player.health = 50;
    player.mana = 500;
    player.usedMana = 0;
    player.shield = 0;
    player.poison = 0;
    player.recharge = 0;
    
    bossStruct boss;
    boss.health = 71;
    boss.damage = 10;
    
    cout << turn(player, boss) << endl;
}

