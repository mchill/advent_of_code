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
    int stats[9][3];
    int current[9][4];
    
    for(int i = 0; i < lines.size(); i++) {
        istringstream iss(lines.at(i));
        string garbage;
        int speed, fly, rest;

        iss >> garbage >> garbage >> garbage >> speed >> garbage >> garbage >>
            fly >> garbage >> garbage >> garbage >> garbage >> garbage >>
            garbage >> rest >> garbage;
            
        stats[i][0] = speed;    // flying speed
        stats[i][1] = fly;      // flying time
        stats[i][2] = rest;     // resting time
        
        current[i][0] = 1;      // mode (resting = 0, flying = 1)
        current[i][1] = fly;    // time left for mode
        current[i][2] = 0;      // distance traveled
        current[i][3] = 0;      // points
    }
    
    for (int second = 0; second < totalTime; second++) {
        int recordDistance = 0;
        
        for (int i = 0; i < lines.size(); i++) {
            if (current[i][1] == 0) {
                current[i][1] = stats[i][1 + current[i][0]];
                current[i][0] = (current[i][0] + 1) % 2;
            }
            
            if (current[i][0] == 1)
                current[i][2] += stats[i][0];
            
            current[i][1]--;
            
            recordDistance = max(recordDistance, current[i][2]);
        }
        
        for (int i = 0; i < lines.size(); i++)
            if (current[i][2] == recordDistance)
                current[i][3]++;
    }
    
    int record = 0;
    for (int i = 0; i < lines.size(); i++)
        record = max(record, current[i][3]);
    
    cout << record << endl;
}

