#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <climits>
#include <algorithm>
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

vector<vector<vector<int> > > getCombinations(int n, int r) {
    vector<vector<vector<int> > > combinations;
    
    vector<bool> v(n);
    fill(v.begin() + n - r, v.end(), true);

    do {
        vector<vector<int> > combinationPair;
        vector<int> combination;
        vector<int> nonCombination;
        
        for (int i = 0; i < n; ++i) {
            if (v[i])
                combination.push_back(i);
            else
                nonCombination.push_back(i);
        }
        
        combinationPair.push_back(combination);
        combinationPair.push_back(nonCombination);
        
        combinations.push_back(combinationPair);
    } while (next_permutation(v.begin(), v.end()));
    
    return combinations;
}

int main () {
    vector<string> lines = getLines();
    vector<int> packages;
    int total = 0;
    
    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        int size = atoi((*line).c_str());
        packages.push_back(size);
        total += size;
    }
    
    int third = total / 3;
    bool match = false;
    unsigned long long entanglement = LLONG_MAX;
    
    for (int r = 1; r < packages.size() / 3; r++) {
        vector<vector<vector<int> > > combinations = getCombinations(packages.size(), r);
        
        for(vector<vector<vector<int> > >::iterator combinationPair = combinations.begin(); combinationPair != combinations.end(); ++combinationPair) {
            vector<int> combination = (*combinationPair).at(0);
            vector<int> nonCombination = (*combinationPair).at(1);
            
            int weight = 0;
            unsigned long long curEntanglement = 1;
            
            for(vector<int>::iterator package = combination.begin(); package != combination.end(); ++package) {
                weight += packages[*package];
                curEntanglement *= packages[*package];
                
                if (weight > third) break;
            }
            
            if (weight != third) continue;
                
            match = true;
            if (curEntanglement > 0 && curEntanglement < entanglement)
                entanglement = curEntanglement;
        }
        
        if (match) break;
    }
    
    cout << entanglement << endl;
}

