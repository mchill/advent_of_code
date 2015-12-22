#include <iostream>
#include <string>
#include <stdlib.h>
#include "md5.h"

using namespace std;

int main () {
    string input = "ckczppom";
    string output = "garbage";
    
    int i = 1;
    for (i; output.find("000000") != 0; i++)
        output = md5(input + to_string(i));
    
    cout << i - 1 << endl;
}

