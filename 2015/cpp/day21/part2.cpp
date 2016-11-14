#include <iostream>
#include <stdlib.h>

using namespace std;

int main () {
    int health = 100;
    int bossHealth = 104;
    int bossDamage = 8;
    int bossArmor = 1;

    int weapons[][5] = {
        {8, 4},
        {10, 5},
        {25, 6},
        {40, 7},
        {74, 8}
    };
    int armors [][6] = {
        {0, 0},
        {13, 1},
        {31, 2},
        {53, 3},
        {75, 4},
        {102, 5}
    };
    int rings[][8] = {
        {0, 0, 0},
        {0, 0, 0},
        {25, 1, 0},
        {50, 2, 0},
        {100, 3, 0},
        {20, 0, 1},
        {40, 0, 2},
        {80, 0, 3}
    };

    int record = 0;

    for (int weapon = 0; weapon < 5; weapon++) {
        for (int armor = 0; armor < 6; armor++) {
            for (int ring1 = 0; ring1 < 8; ring1++) {
                for (int ring2 = 0; ring2 < 8; ring2++) {
                    if (ring1 == ring2) continue;

                    int gold = weapons[weapon][0] + armors[armor][0] +
                                rings[ring1][0] + rings[ring2][0];

                    if (gold < record) continue;

                    int damage = weapons[weapon][1] + rings[ring1][1] + rings[ring2][1];
                    int defense = armors[armor][1] + rings[ring1][2] + rings[ring2][2];

                    int hitsToKillBoss = (bossHealth - 1) / max(damage - bossArmor, 1);
                    int hitsToKillPlayer = (health - 1) / max(bossDamage - defense, 1);

                    if (hitsToKillPlayer < hitsToKillBoss)
                        record = max(record, gold);
                }
            }
        }
    }

    cout << record << endl;
}
