package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	DIRECTIONS := [][]int{
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 0},
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var diagram []string

	for scanner.Scan() {
		diagram = append(diagram, scanner.Text())
	}

	var encountered string
	var steps int
	dir := 0
	x := strings.IndexByte(diagram[0], '|')
	y := 0

	for steps = 0; diagram[y][x] != ' '; steps++ {
		switch diagram[y][x] {
		case '+':
			dir = (dir + 1) % len(DIRECTIONS)
			if diagram[y+DIRECTIONS[dir][1]][x+DIRECTIONS[dir][0]] == ' ' {
				dir = (dir - 2) % len(DIRECTIONS)
				if dir < 0 {
					dir += len(DIRECTIONS)
				}
			}
		case '|':
		case '-':
		default:
			encountered += string(diagram[y][x])
		}

		x += DIRECTIONS[dir][0]
		y += DIRECTIONS[dir][1]
	}

	fmt.Println("Part 1:", encountered)
	fmt.Println("Part 2:", steps)
}
