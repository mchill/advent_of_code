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

string getChars() {
    ifstream file("input.txt");
    string input = "";
    char c;

    while (file.get(c))
        input += c;

    return input;
}

void find_and_replace(string& source, string const& find, string const& replace) {
    for(string::size_type i = 0; (i = source.find(find, i)) != string::npos;)
    {
        source.replace(i, find.length(), replace);
        i += replace.length();
    }
}

int main () {
    int total = 0;

    vector<string> lines = getLines();

    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        find_and_replace(*line, "x", " ");

        istringstream iss(*line);
        int length, width, height;

        iss >> length >> width >> height;

        total += 2 * min(min(length + width, length + height), width + height) + length * width * height;
    }

    cout << total << endl;
}

