#include <iostream>
#include <math.h>
#include <stdlib.h>

using namespace std;

int main () {
	int minimum = 33100000;
	int gifts = 0;
	int house = 1;

	while (gifts < minimum) {
		gifts = 0;

		int squareRoot = (int) sqrt(house) + 1;
		for (int i = 1; i < squareRoot; i++) {
			if (house % i == 0) {
				gifts += 10 * i;
				if (house / i != i)
					gifts += 10 * house / i;
			}
		}

		house++;
	}

	cout << house - 1 << endl;
}
