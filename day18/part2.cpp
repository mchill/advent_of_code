#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <stdlib.h>

using namespace std;

const int SIZE = 100;

vector<string> getLines() {
    ifstream file("input.txt");
    vector<string> lines;
    string line;

    while (getline(file, line))
        lines.push_back(line);

    return lines;
}

int countNeighbors(bool *cell) {
    return *(cell + 1) + *(cell + SIZE + 3) + *(cell + SIZE + 2) + *(cell + SIZE + 1) +
            *(cell - 1) + *(cell - SIZE - 3) + *(cell - SIZE - 2) + *(cell - SIZE - 1);
}

int main () {
    vector<string> lines = getLines();
    bool grid1[SIZE + 2][SIZE + 2] = { false };
    bool grid2[SIZE + 2][SIZE + 2] = { false };
    
    for(int y = 0; y < SIZE; y++)
        for (int x = 0; x < SIZE; x++)
            if (lines.at(y)[x] == '#')
                grid1[x+1][y+1] = grid2[x+1][y+1] = true;
            
    grid1[1][1] = grid2[1][1] = true;
    grid1[1][SIZE] = grid2[1][SIZE] = true;
    grid1[SIZE][1] = grid2[SIZE][1] = true;
    grid1[SIZE][SIZE] = grid2[SIZE][SIZE] = true;
    
    for (int step = 0; step < 100; step++) {
        for(int y = 1; y <= SIZE; y++) {
            for (int x = 1; x <= SIZE; x++) {
                int neighbors = countNeighbors(&(grid1[x][y]));
                
                if (grid1[x][y] && !(neighbors == 2 || neighbors == 3) &&
                    !((x == 1 && y == 1) || (x == 1 && y == SIZE) ||
                      (x == SIZE && y == 1) || (x == SIZE && y == SIZE)))
                    grid2[x][y] = false;
                else if (!grid1[x][y] && neighbors == 3)
                    grid2[x][y] = true;
            }
        }
        
        for(int y = 1; y <= SIZE; y++)
            for (int x = 1; x <= SIZE; x++)
                grid1[x][y] = grid2[x][y];
    }
    
    int on = 0;
    
    for(int y = 1; y <= SIZE; y++)
        for (int x = 1; x <= SIZE; x++)
            on += grid1[x][y];
        
    cout << on << endl;
}

