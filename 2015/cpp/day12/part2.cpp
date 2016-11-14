#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <stack>
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
    string file = getChars();

    stack<int> lastOpen;
    int red = 0;
	
    for (int i = 0; i < file.size(); i++) {
        if (red) {
            if (file[i] == '{')
                red++;
            else if (file[i] == '}')
                red--;
            
            if (!red) {
                file.erase(file.begin()+lastOpen.top(), file.begin()+i+1);
                i = lastOpen.top();
                lastOpen.pop();
            }
			
            continue;
        }
		
        if (file[i] == '{')
            lastOpen.push(i);
        else if (file[i] == '}')
            lastOpen.pop();
        else if (file[i] == ':' &&
          file[i+1] == '"' &&
          file[i+2] == 'r' &&
          file[i+3] == 'e' &&
          file[i+4] == 'd' &&
          file[i+5] == '"')
            red++;
    }
    
    int total = 0;
	int value;
	string str;
	
	replace(file.begin(), file.end(), ':', ' ');
	replace(file.begin(), file.end(), ',', ' ');
	replace(file.begin(), file.end(), '[', ' ');
	replace(file.begin(), file.end(), ']', ' ');
	replace(file.begin(), file.end(), '}', ' ');
	
	istringstream iss(file);
	
	while (iss >> str) {
		try {
			value = stoi(str);
			total += value;
		} catch (invalid_argument&) {
			continue;
		}
	}
	
	cout << total << endl;
}

