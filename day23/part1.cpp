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

int main () {
    vector<string> lines = getLines();
    vector<vector<int> > instructions;
    int registers[2] = {0};
    
    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        replace((*line).begin(), (*line).end(), ',', ' ');
        replace((*line).begin(), (*line).end(), '+', ' ');
        
        istringstream iss1(*line), iss2(*line), iss3(*line);
        string instruction;
        char reg = 0;
        int jump = 0;

        if (!(iss1 >> instruction >> jump))
            iss2 >> instruction >> reg >> jump;
        
        int instructionIndex;
        if (instruction == "hlf") instructionIndex = 0;
        else if (instruction == "tpl") instructionIndex = 1;
        else if (instruction == "inc") instructionIndex = 2;
        else if (instruction == "jmp") instructionIndex = 3;
        else if (instruction == "jie") instructionIndex = 4;
        else if (instruction == "jio") instructionIndex = 5;
        
        instructions.push_back(vector<int> {instructionIndex, reg, jump});
    }
    
    for (int i = 0; i < instructions.size(); ) {
        char reg = instructions[i].at(1);
        
        switch (instructions[i].at(0)) {
            case 0:
                registers[reg - 'a'] /= 2;
                i++;
                break;
            case 1:
                registers[reg - 'a'] *= 3;
                i++;
                break;
            case 2:
                registers[reg - 'a']++;
                i++;
                break;
            case 3:
                i += instructions[i].at(2);
                break;
            case 4:
                if (registers[reg - 'a'] % 2 == 0)
                    i += instructions[i].at(2) - 1;
                i++;
                break;
            case 5:
                if (registers[reg - 'a'] == 1)
                    i += instructions[i].at(2) - 1;
                i++;
        }
    }
    
    cout << registers[1] << endl;
}
