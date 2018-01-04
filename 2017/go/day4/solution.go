package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(word string) string {
	slice := []rune(word)
	sort.Slice(slice, func(index1, index2 int) bool {
		return slice[index1] < slice[index2]
	})
	return string(slice)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		wordMap := make(map[string]bool)
		sortedWordMap := make(map[string]bool)
		foundAnagram := false
		words := strings.Split(scanner.Text(), " ")
		part1++

		for _, word := range words {
			sortedWord := sortString(word)

			if wordMap[word] {
				part1--
				foundAnagram = true
				break
			} else if !foundAnagram && sortedWordMap[sortedWord] {
				foundAnagram = true
			}

			wordMap[word] = true
			sortedWordMap[sortedWord] = true
		}

		if !foundAnagram {
			part2++
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
