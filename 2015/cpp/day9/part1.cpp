#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <map>
#include <queue>
#include <stdlib.h>

using namespace std;

int distances[8][8] = {0};

vector<string> getLines() {
    ifstream file("input.txt");
    vector<string> lines;
    string line;

    while (getline(file, line))
        lines.push_back(line);

    return lines;
}

string getChars() {
    ifstream file("input.txt");
    string input = "";
    char c;

    while (file.get(c))
        input += c;

    return input;
}

void findAndReplace(string& source, string const& find, string const& replace) {
    for(string::size_type i = 0; (i = source.find(find, i)) != string::npos;)
    {
        source.replace(i, find.length(), replace);
        i += replace.length();
    }
}

void printLines(vector<string> lines) {
    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        cout << *line << endl;
    }
}

void printChars(string line) {
    for (string::iterator c = line.begin(); c != line.end(); ++c) {
        cout << *c << endl;
    }
}

void istringstreamExample(string line) {
    istringstream iss(line);
    string myString;
    int myInt;

    if (!(iss >> myString >> myInt)) {
        cout << "Error!\n";
    }
}

int travel(queue<int> choices, int last = -1) {
    int smallest = 0;

    for (int i = 0; i < choices.size(); i++) {
        int next = choices.front();
        choices.pop();

        if (last == -1)
            smallest = travel(choices, next);
        else if (smallest == 0)
            smallest = travel(choices, next) + distances[last][next];
        else
            smallest = min(smallest, travel(choices, next) + distances[last][next]);

        choices.push(next);
    }

    return smallest;
}

int main () {
    vector<string> lines = getLines();
    map<string, int> locations = {
        {"AlphaCentauri", 0},
        {"Snowdin", 1},
        {"Tambi", 2},
        {"Faerun", 3},
        {"Norrath", 4},
        {"Straylight", 5},
        {"Tristram", 6},
        {"Arbre", 7}
    };
    queue<int> choices({0, 1, 2, 3, 4, 5, 6, 7});
    int smallest = 9000;

    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        istringstream iss(*line);
        string origin, destination, garbage;
        int distance;

        if (!(iss >> origin >> garbage >> destination >> garbage >> distance)) {
            cout << "Error!\n";
            exit(1);
        }

        distances[locations[origin]][locations[destination]] = distance;
        distances[locations[destination]][locations[origin]] = distance;
    }

    cout << travel(choices) << endl;
}

