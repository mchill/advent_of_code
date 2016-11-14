#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <stdlib.h>

using namespace std;

int lights[1000][1000];

void on(int x1, int y1, int x2, int y2) {
    for (int i = x1; i <= x2; i++) {
        for (int j = y1; j <= y2; j++) {
            lights[i][j] = 1;
        }
    }
}

void off(int x1, int y1, int x2, int y2) {
    for (int i = x1; i <= x2; i++) {
        for (int j = y1; j <= y2; j++) {
            lights[i][j] = 0;
        }
    }
}

void toggle(int x1, int y1, int x2, int y2) {
    for (int i = x1; i <= x2; i++) {
        for (int j = y1; j <= y2; j++) {
            lights[i][j] = (lights[i][j] + 1) % 2;
        }
    }
}

int count() {
    int count = 0;

    for (int i = 0; i < 1000; i++) {
        for (int j = 0; j < 1000; j++) {
            count += lights[i][j];
        }
    }

    return count;
}

void find_and_replace(string& source, string const& find, string const& replace) {
    for(string::size_type i = 0; (i = source.find(find, i)) != string::npos;)
    {
        source.replace(i, find.length(), replace);
        i += replace.length();
    }
}

int main ()
{
    ifstream infile("input.txt");

    string line;
    while (getline(infile, line))
    {
        find_and_replace(line, ",", " ");

        istringstream iss(line), iss2(line);
        string garbage, action;
        int x1, x2, y1, y2;

        if (!(iss >> garbage >> action >> x1 >> y1 >> garbage >> x2 >> y2) &&
            !(iss2 >> action >> x1 >> y1 >> garbage >> x2 >> y2)) {
                cout << "Error!\n";
        }

        if (action == "on") {
            on(x1, y1, x2, y2);
        } else if (action == "off") {
            off(x1, y1, x2, y2);
        } else if (action == "toggle") {
            toggle(x1, y1, x2, y2);
        } else {
            cout << "Error!\n";
        }
    }

    cout << count() << "\n";
}