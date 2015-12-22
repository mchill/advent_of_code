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

int main () {
	string file = getChars();
	
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

