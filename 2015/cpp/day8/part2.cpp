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

int count(string& source, string const& find) {
    int count = 0;
    for(string::size_type i = 0; (i = source.find(find, i)) != string::npos;)
    {
        count++;
        i += find.length();
    }
    return count;
}

int main () {
    vector<string> lines = getLines();
    int diff = 0;

    for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        string str = (*line).substr(1, (*line).size() - 2);
        
        int numSlashQuote = count(str, "\\\"");
        int numSlashSlash = count(str, "\\\\");
        int numSlashX = count(str, "\\x");
        int numSlashSlashX = count(str, "\\\\x");
        int numSlashSlashSlashX = count(str, "\\\\\\x");
        int numHexChars = numSlashX - numSlashSlashX + numSlashSlashSlashX;

        diff += numSlashQuote * 2 + numSlashSlash * 2 + numHexChars + 4;
    }

    cout << diff << endl;
}

