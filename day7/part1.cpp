#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <map>
#include <stdlib.h>
#include <cstdlib>

using namespace std;

map<string, string> lineMap;
map<string, int> solved;

vector<string> getLines() {
    ifstream file("input.txt");
    vector<string> lines;
    string line;

    while (getline(file, line))
        lines.push_back(line);

    return lines;
}

unsigned short calculate(string wire) {
    if (solved.find(wire) != solved.end()) return solved[wire];

    string line = lineMap[wire];

    istringstream iss1(line), iss2(line), iss3(line);
    string arg1, arg2, result, operation, garbage;
    char *p;

    if (iss1 >> arg1 >> operation >> arg2 >> garbage >> result) {
        unsigned short convertedArg1 = strtol(arg1.c_str(), &p, 10);
        if (*p) convertedArg1 = calculate(arg1);

        unsigned short convertedArg2 = strtol(arg2.c_str(), &p, 10);
        if (*p) convertedArg2 = calculate(arg2);

        if (operation == "AND") solved[wire] = convertedArg1 & convertedArg2;
        else if (operation == "OR") solved[wire] = convertedArg1 | convertedArg2;
        else if (operation == "LSHIFT") solved[wire] = convertedArg1 << convertedArg2;
        else if (operation == "RSHIFT") solved[wire] = convertedArg1 >> convertedArg2;

        return solved[wire];
    } else if (iss2 >> operation >> arg1 >> garbage >> result) {
        unsigned short converted = strtol(arg1.c_str(), &p, 10);
        if (*p) converted = calculate(arg1);

        solved[wire] = ~converted;
        return solved[wire];
    } else if (iss3 >> arg1 >> operation >> result) {
        unsigned short converted = strtol(arg1.c_str(), &p, 10);
        if (*p) converted = calculate(arg1);

        solved[wire] = converted;
        return solved[wire];
    } else {
        cout << "Error!" << endl;
        exit(1);
    }
}

int main () {
    vector<string> lines = getLines();

    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        string result = (*line).substr((*line).rfind("->") + 3);
        lineMap[result] = *line;
    }

    cout << calculate("a") << endl;
}

