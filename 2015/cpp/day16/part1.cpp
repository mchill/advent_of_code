#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <algorithm>
#include <vector>
#include <map>
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
    map<string, int> evidence = {
        {"children", 3},
        {"cats", 7},
        {"samoyeds", 2},
        {"pomeranians", 3},
        {"akitas", 0},
        {"vizslas", 0},
        {"goldfish", 5},
        {"trees", 3},
        {"cars", 2},
        {"perfumes", 1}
    };
    
    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        replace((*line).begin(), (*line).end(), ',', ' ');
        replace((*line).begin(), (*line).end(), ':', ' ');
        
        istringstream iss(*line);
        string garbage, category1, category2, category3;
        int sue, value1, value2, value3;

        iss >> garbage >> sue >> category1 >> value1
            >> category2 >> value2 >> category3 >> value3;
            
        if (evidence[category1] == value1 &&
            evidence[category2] == value2 &&
            evidence[category3] == value3) {
            cout << sue << endl;
            break;
        }
    }
}

