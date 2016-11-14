#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>
#include <stdlib.h>

using namespace std;

string getChars() {
    ifstream file("input.txt");
    string input = "";
    char c;

    while (file.get(c))
        input += c;

    return input;
}

int main () {
    string directions = getChars();
    vector<string> visited = {"0,0"};
    int *position = NULL;
    int position1[2] = {0, 0};
    int position2[2] = {0, 0};

    bool santa = true;

    for (string::iterator c = directions.begin(); c != directions.end(); ++c) {
        if (santa) position = position1;
        else position = position2;
        santa = !santa;

        if (*c == '>') position[0]++;
        else if (*c == '<') position[0]--;
        else if (*c == 'v') position[1]++;
        else if (*c == '^') position[1]--;

        string location = to_string(position[0]) + "," + to_string(position[1]);

        if (find(visited.begin(), visited.end(), location) == visited.end()) {
            visited.push_back(location);
        }
    }

    cout << visited.size() << endl;
}

