package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	registers := make(map[string]int)
	var part1 int
	var part2 int

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		register := words[0]
		operation := 1
		if words[1] == "dec" {
			operation = -1
		}
		offset, _ := strconv.Atoi(words[2])
		conditionalRegister := words[4]
		conditionalOperator := words[5]
		conditionalValue, _ := strconv.Atoi(words[6])

		var conditional bool
		switch conditionalOperator {
		case "==":
			conditional = registers[conditionalRegister] == conditionalValue
		case "!=":
			conditional = registers[conditionalRegister] != conditionalValue
		case ">":
			conditional = registers[conditionalRegister] > conditionalValue
		case ">=":
			conditional = registers[conditionalRegister] >= conditionalValue
		case "<":
			conditional = registers[conditionalRegister] < conditionalValue
		case "<=":
			conditional = registers[conditionalRegister] <= conditionalValue
		}

		if conditional {
			registers[register] += operation * offset
			if registers[register] > part2 {
				part2 = registers[register]
			}
		}
	}

	for _, value := range registers {
		if value > part1 {
			part1 = value
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
