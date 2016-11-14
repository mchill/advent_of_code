#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
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

int main () {
    vector<string> lines = getLines();
    const int totalTime = 2503;
    int record = 0;
    
    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        int time = 0;
        int distance = 0;
        bool flying = false;
        
        istringstream iss(*line);
        string garbage;
        int speed, fly, rest;

        iss >> garbage >> garbage >> garbage >> speed >> garbage >> garbage >>
            fly >> garbage >> garbage >> garbage >> garbage >> garbage >>
            garbage >> rest >> garbage;
            
        while (time < totalTime) {
            flying = !flying;
            
            if (flying) {
                distance += speed * min(fly, totalTime - time);
                time += fly;
                continue;
            }
            
            time += rest;
        }
        
        record = max(record, distance);
    }
    
    cout << record << endl;
}

