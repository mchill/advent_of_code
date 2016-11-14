#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <algorithm>
#include <vector>
#include <stdlib.h>

using namespace std;

vector<string> getLines() {
    ifstream file("input.txt");
    vector<string> lines;
    string line;

    while (getline(file, line))
        lines.push_back(line);

    return lines;
}

void istringstreamExample(string line) {

}

int main () {
    vector<string> lines = getLines();
    int ingredients[4][4];

    for(int i = 0; i < lines.size(); i++) {
        replace(lines.at(i).begin(), lines.at(i).end(), ',', ' ');

        istringstream iss(lines.at(i));
        string garbage;

        iss >> garbage >> garbage >> ingredients[i][0] >> garbage >> ingredients[i][1]
            >> garbage >> ingredients[i][2] >> garbage >> ingredients[i][3];
    }

    int record = 0;

    for (int sugar = 0; sugar < 100; sugar++) {
        for (int sprinkles = 0; sprinkles < 100 - sugar; sprinkles++) {
            for (int candy = 0; candy < 100 - sugar - sprinkles; candy++) {
                int chocolate = 100 - sugar - sprinkles - candy;

                int total = 1;

                for (int ingredient = 0; ingredient < 4; ingredient++) {
                    int subTotal = ingredients[0][ingredient] * sugar +
                                   ingredients[1][ingredient] * sprinkles +
                                   ingredients[2][ingredient] * candy +
                                   ingredients[3][ingredient] * chocolate;

                    if (subTotal <= 0) {
                        total = 0;
                        break;
                    }

                    total *= subTotal;
                }

                record = max(record, total);
            }
        }
    }

    cout << record << endl;
}

