#include <iostream>
#include <string>
#include <string.h>
#include <stdlib.h>

using namespace std;

int main () {
	char password[] = "hxbxxyzz";
	const int length = strlen(password);
	int index = length - 1;

	bool straightRule = false;
	bool pairRule = false;

	while (!straightRule || !pairRule) {
		straightRule = false;
		pairRule = false;

		if (password[index] == 'z') {
			password[index] = 'a';
			index--;
			continue;
		}

		password[index]++;

		char cur = password[index];
		if (cur == 'i' || cur == 'o' || cur == 'l') {
			password[index]++;

			for (int i = index + 1; i < length; i++)
				password[i] = 'a';
		}

		index = length - 1;

		char last = ' ';
		int straight = 0;
		char match = '\0';

		for (int i = 0; i < length; i++) {
			if (last + 1 == password[i])
				straight++;
			else
				straight = 0;
			if (straight == 2)
				straightRule = true;

			if (last == password[i]) {
				if (!match)
					match = password[i];
				else if (match != password[i])
					pairRule = true;
			}

			last = password[i];
		}
	}

	cout << password << endl;
}

