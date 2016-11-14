#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <stdlib.h>

using namespace std;

int main ()
{
    ifstream infile("input.txt");
    int count = 0;

    string line;
    while (getline(infile, line))
    {
        char last = ' ';
        char twoago = ' ';
        vector<string> pairs;
        
        bool rule1 = false;
        bool rule2 = false;

        for (string::iterator c = line.begin(); c != line.end(); ++c) {
            if (!rule1 && *c == twoago)
                rule1 = true;

            if (!rule2) {
                string pair = "";
                pair += last;
                pair += *c;

                for (int i = 1; i < pairs.size(); ++i)
                    if (pair == pairs.at(i-1))
                        rule2 = true;

                pairs.push_back(pair);
            }

            twoago = last;
            last = *c;
        }

        if (rule1 && rule2)
            count++;
    }

    cout << count << endl;
}
