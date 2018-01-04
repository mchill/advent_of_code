package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func abs(number int) int {
	if number < 0 {
		number = -number
	}
	return number
}

func min(number1, number2 int) int {
	if number1 < number2 {
		return number1
	}
	return number2
}

func max(number1, number2 int) int {
	if number1 > number2 {
		return number1
	}
	return number2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	directions := strings.Split(line, ",")

	xPos, yPos, distance := 0, 0, 0

	for _, direction := range directions {
		switch direction {
		case "n":
			yPos--
		case "ne":
			xPos++
		case "se":
			xPos++
			yPos++
		case "s":
			yPos++
		case "sw":
			xPos--
		case "nw":
			xPos--
			yPos--
		}
		distance = max(distance, max(abs(xPos-yPos), max(xPos, yPos)))
	}

	fmt.Println("Part 1:", max(abs(xPos-yPos), max(xPos, yPos)))
	fmt.Println("Part 2:", distance)
}
