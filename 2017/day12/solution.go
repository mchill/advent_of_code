package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeGroup(program int, programs map[int][]int) {
	reachablePrograms := programs[program]
	delete(programs, program)

	for _, reachable := range reachablePrograms {
		removeGroup(reachable, programs)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	programs := make(map[int][]int)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		program, _ := strconv.Atoi(words[0])

		for _, word := range words[2:] {
			reachable, _ := strconv.Atoi(strings.TrimRight(word, ","))
			programs[program] = append(programs[program], reachable)
		}
	}

	groups := 1
	numPrograms := len(programs)
	removeGroup(0, programs)

	fmt.Println("Part 1:", numPrograms-len(programs))

	for len(programs) > 0 {
		var program int
		for program, _ = range programs {
			break
		}
		groups++
		removeGroup(program, programs)
	}

	fmt.Println("Part 2:", groups)

}
