package main

import (
	"bufio"
	"fmt"
	"os"
)

var DIRECTIONS = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func part1(cluster map[string]bool, row int) int {
	curX := row / 2
	curY := row / 2
	dir := 0
	infected := 0

	for burst := 0; burst < 10000; burst++ {
		current := fmt.Sprintf("%d,%d", curX, curY)
		dir = (dir + 1) % len(DIRECTIONS)

		if !cluster[current] {
			dir = (dir + 2) % len(DIRECTIONS)
			infected++
		}

		cluster[current] = !cluster[current]
		curX += DIRECTIONS[dir][0]
		curY += DIRECTIONS[dir][1]
	}

	return infected
}

func part2(cluster map[string]int, row int) int {
	curX := row / 2
	curY := row / 2
	dir := 0
	infected := 0

	for burst := 0; burst < 10000000; burst++ {
		current := fmt.Sprintf("%d,%d", curX, curY)
		dir = (dir + cluster[current] + 3) % len(DIRECTIONS)

		if cluster[current] == 1 {
			infected++
		}

		cluster[current] = (cluster[current] + 1) % 4
		curX += DIRECTIONS[dir][0]
		curY += DIRECTIONS[dir][1]
	}

	return infected
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	cluster1 := make(map[string]bool)
	cluster2 := make(map[string]int)
	var row int

	for scanner.Scan() {
		for col, node := range scanner.Text() {
			if node == '#' {
				current := fmt.Sprintf("%d,%d", col, row)
				cluster1[current] = true
				cluster2[current] = 2
			}
		}
		row++
	}

	fmt.Println("Part 1:", part1(cluster1, row))
	fmt.Println("Part 2:", part2(cluster2, row))
}
