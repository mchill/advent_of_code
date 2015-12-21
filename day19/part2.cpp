#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <string.h>
#include <vector>
#include <map>
#include <algorithm>
#include <stdlib.h>

using namespace std;

int main () {
	/*
	I stared at the input for a while, and discovered a pattern. Every reaction
	follows one of a few rules:

		e -> XX
		X -> XX

		X -> X Rn X Ar
		X -> X Rn X Y X Ar
		X -> X Rn X Y X Y X Ar

	Using that information, I constructed a simple formula to determine how
	many steps would transform e to the final molecule:

	Elements           | Steps
	-------------------|----------------------------
	Total              | 0
	Total-3*(Rn-Y)     | Rn-Y
	Total-3*(Rn-Y)-5*Y | Rn-Y+Y
	1                  | Rn-Y+Y+Total-3*(Rn-Y)-5*Y-1
	1                  | Rn+Total-3*Rn+3*Y-5*Y-1
	1                  | Total-2*Rn-2*Y-1
	*/

	string molecule = "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr";
	int total = 0;
	int Rn = 0;
	int Y = 0;

	for (string::iterator c = molecule.begin(); c != molecule.end(); ++c) {
        if (*c >= 'A' && *c <= 'Z')
			total++;
		if (*c == 'R')
			Rn++;
		if (*c == 'Y')
			Y++;
    }

	cout << total - 2 * (Rn + Y) - 1 << endl;
}
