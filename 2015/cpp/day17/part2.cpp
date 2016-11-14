#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
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

vector<vector<int> > getCombinations(int n) {
    vector<vector<int> > combinations;
    
    for (int r = 1; r <= n; r++) {
        vector<bool> v(n);
        fill(v.begin() + n - r, v.end(), true);

        do {
            vector<int> combination;
            
            for (int i = 0; i < n; ++i) {
                if (v[i]) {
                    combination.push_back(i);
                }
            }
            
            combinations.push_back(combination);
        } while (next_permutation(v.begin(), v.end()));
    }
    
    return combinations;
}

int main () {
    string file = getChars();
    vector<int> containers;
    
    istringstream iss(file);
    int value;

    while (iss >> value)
        containers.push_back(value);
    
    vector<vector<int> > combinations = getCombinations(containers.size());
    int fits = 0;
    int numContainers = containers.size() + 1;
    
    for(vector<vector<int> >::iterator combination = combinations.begin(); combination != combinations.end(); ++combination) {
        int filled = 0;
        
        for(vector<int>::iterator container = (*combination).begin(); container != (*combination).end(); ++container) {
            filled += containers[*container];
            
            if (filled > 150)
                break;
        }
        
        if (filled == 150) {
            if ((*combination).size() > numContainers)
                break;
            else if ((*combination).size() < numContainers) {
                fits = 0;
                numContainers = (*combination).size();
            }

            fits++;
        }
    }
    
    cout << fits << endl;
}
