package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(banks []int) int {
	maximum := 0
	for index := 1; index < len(banks); index++ {
		if banks[index] > banks[maximum] {
			maximum = index
		}
	}
	return maximum
}

func str(banks []int) string {
	bankString := ""
	for _, bank := range banks {
		bankString += string(bank)
	}
	return bankString
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	words := strings.Split(scanner.Text(), "\t")
	var banks []int

	for _, word := range words {
		bank, _ := strconv.Atoi(word)
		banks = append(banks, bank)
	}

	part1 := 0
	seen := make(map[string]int)

	for bankString := ""; seen[bankString] == 0; part1++ {
		seen[bankString] = part1

		largestBank := max(banks)
		value := banks[largestBank]
		banks[largestBank] = 0

		for index := 1; index <= value; index++ {
			banks[(largestBank+index)%len(banks)]++
		}

		bankString = str(banks)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part1-seen[str(banks)])
}
