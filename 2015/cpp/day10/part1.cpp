#include <string>
#include <iostream>
#include <stdlib.h>

using namespace std;

int main () {
    string input = "3113322113";

    for (int i = 0; i < 40; i++) {
        string output = "";

        char last = input.at(0);
        int count = 1;

        for (string::iterator c = input.begin() + 1; c != input.end(); ++c) {
            if (*c != last) {
                output += to_string(count) + last;
                count = 0;
            }
            count++;

            last = *c;
        }

        output += to_string(count) + last;

        input = output;
    }

    cout << input.length() << endl;
}

