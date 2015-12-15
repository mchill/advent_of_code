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
    vector<string> visited;
    int position[2] = {0, 0};

    for (string::iterator c = directions.begin(); c != directions.end(); ++c) {
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

