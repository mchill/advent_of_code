#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <stdlib.h>

using namespace std;

int main ()
{
    ifstream infile("input.txt");
    int count = 0;

    string line;
    while (getline(infile, line))
    {
        int vowels = 0;
        char last = ' ';
        bool match = false;
        bool nope = false;

        for (string::iterator c = line.begin(); c != line.end(); ++c) {
            if (*c == 'a' || *c == 'e' || *c == 'i' || *c == 'o' || *c == 'u') {
                vowels++;
            }

            if (*c == last) {
                match = true;
            }

            if ((last == 'a' && *c == 'b') || (last == 'c' && *c == 'd') || (last == 'p' && *c == 'q') || (last == 'x' && *c == 'y')) {
                nope = true;
                break;
            }

            last = *c;
        }

        if (!nope && match && vowels >= 3) {
            count++;
        }
    }

    cout << count << endl;
}
