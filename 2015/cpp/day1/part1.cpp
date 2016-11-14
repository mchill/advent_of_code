#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <stdlib.h>

using namespace std;

int main ()
{
    ifstream infile("input.txt");
    int floor = 0;

    char c;
    while (infile.get(c))
    {
        if (c == '(')
            floor++;
        else if (c == ')')
            floor--;
    }

    cout << floor << endl;
}
