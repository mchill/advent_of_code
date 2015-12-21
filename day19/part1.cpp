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

vector<string> getLines() {
    ifstream file("input.txt");
    vector<string> lines;
    string line;

    while (getline(file, line))
        lines.push_back(line);

    return lines;
}

int main () {
	vector<string> lines = getLines();
	map<string, vector<string> > equations;

	char first[] = "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr";
	const char *original = first;
	vector<string> results;

	for(vector<string>::iterator line = lines.begin(); line != lines.end(); ++line) {
        istringstream iss(*line);
		string reactant, product;

		iss >> reactant >> product >> product;

		if (equations.find(reactant) != equations.end()) {
			equations[reactant].push_back(product);
			continue;
		}

		equations[reactant] = vector<string> { product };
    }

	for (int i = 0; i < strlen(original); ) {
		string next = "";
		next += original[i];

		if (equations.find(next) != equations.end()) {
			vector<string> products = equations[next];

			for (vector<string>::iterator product = products.begin(); product != products.end(); ++product) {
				string result = first;
				result.replace(i, 1, *product);
				
				if (find(results.begin(), results.end(), result) == results.end())
					results.push_back(result);
			}
			i++;
			continue;
		}

		i++;
		next += original[i];

		if (equations.find(next) != equations.end()) {
			vector<string> products = equations[next];

			for (vector<string>::iterator product = products.begin(); product != products.end(); ++product) {
				string result = first;
				result.replace(i-1, 2, *product);
				
				if (find(results.begin(), results.end(), result) == results.end())
					results.push_back(result);
			}
			i++;
		}
	}

	cout << results.size() << endl;
}
