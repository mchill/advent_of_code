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
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		var numbers []int
		words := strings.Split(scanner.Text(), "\t")
		min, _ := strconv.Atoi(words[0])
		max, _ := strconv.Atoi(words[1])

		for _, word := range words {
			number, _ := strconv.Atoi(word)
			numbers = append(numbers, number)

			if number < min {
				min = number
			} else if number > max {
				max = number
			}
		}

		part1 += max - min

		for index1, number1 := range numbers {
			for index2, number2 := range numbers {
				if index1 == index2 {
					continue
				} else if number1%number2 == 0 {
					part2 += number1 / number2
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
