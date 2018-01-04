package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	part1 := 0
	part2 := 0
	level := 0
	garbage := false
	cancelled := false

	for _, char := range input {
		if cancelled {
			cancelled = false
		} else if char == '!' {
			cancelled = true
		} else if char == '>' {
			garbage = false
		} else if garbage {
			part2++
			continue
		} else if char == '<' {
			garbage = true
		} else if char == '{' {
			level++
			part1 += level
		} else if char == '}' {
			level--
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
