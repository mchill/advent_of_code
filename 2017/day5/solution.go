package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var lines1 []int
	var lines2 []int
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines1 = append(lines1, line)
		lines2 = append(lines2, line)
	}

	line := 0
	for line >= 0 && line < len(lines1) {
		lines1[line]++
		line += lines1[line] - 1
		part1++
	}

	line = 0
	for line >= 0 && line < len(lines2) {
		offset := -1
		if lines2[line] < 3 {
			offset = 1
		}
		lines2[line] += offset
		line += lines2[line] - offset
		part2++
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
