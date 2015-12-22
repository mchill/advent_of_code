#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <map>
#include <queue>
#include <stdlib.h>

using namespace std;

int data[8][8] = {0};

vector<string> getLines() {
    ifstream file("input.txt");
    vector<string> lines;
    string line;

    while (getline(file, line))
        lines.push_back(line);

    return lines;
}

int nextPerson(queue<int> choices, int last) {
    int biggest = 0;
    
    for (int i = 0; i < choices.size(); i++) {
        int next = choices.front();
        choices.pop();

        int happiness = data[last][next] + data[next][last];
        happiness += nextPerson(choices, next);
        biggest = max(biggest, happiness);

        choices.push(next);
    }
    
    if (choices.size() == 0)
        biggest += data[last][0] + data[0][last];
    
    return biggest;
}

int main () {
    vector<string> lines = getLines();
    
    map<string, int> people = {
        {"Alice", 0},
        {"Bob", 1},
        {"Carol", 2},
        {"David", 3},
        {"Eric", 4},
        {"Frank", 5},
        {"George", 6},
        {"Mallory", 7}
    };
    queue<int> choices({1, 2, 3, 4, 5, 6, 7});
    
    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        istringstream iss(*line);
        string person, otherPerson, action, garbage;
        int happiness;
        
        iss >> person >> garbage >> action >> happiness >> garbage >> garbage >> garbage >> garbage >> garbage >> garbage >> otherPerson;
        otherPerson.pop_back();
        
        if (action == "lose") happiness *= -1;
        data[people[person]][people[otherPerson]] = happiness;
    }
    
    cout << nextPerson(choices, 0) << endl;
}

