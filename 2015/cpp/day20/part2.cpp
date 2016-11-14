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
				int pair = house / i;
				if (pair < 50)
					gifts += 11 * i;
				if (pair != i && house / pair < 50)
					gifts += 11 * pair;
			}
		}

		house++;
	}

	cout << house - 1 << endl;
}
