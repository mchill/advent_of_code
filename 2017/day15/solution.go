package main

import (
	"fmt"
)

const FACTORA = 16807
const FACTORB = 48271
const DIVIDE = 2147483647

type Matches struct {
	compared int
	matching int
	matchesA []int64
	matchesB []int64
}

func generate(genA, genB int64, matches *Matches) (int64, int64) {
	genA = genA * FACTORA % DIVIDE
	genB = genB * FACTORB % DIVIDE

	if genA%4 == 0 {
		matches.matchesA = append(matches.matchesA, genA)
	}
	if genB%8 == 0 {
		matches.matchesB = append(matches.matchesB, genB)
	}

	if len(matches.matchesA) > 0 && len(matches.matchesB) > 0 {
		matches.compared++

		var matchA, matchB int64
		matches.matchesA, matchA = matches.matchesA[1:], matches.matchesA[0]
		matches.matchesB, matchB = matches.matchesB[1:], matches.matchesB[0]

		if matchA&0xFFFF == matchB&0xFFFF {
			matches.matching++
		}
	}

	return genA, genB
}

func main() {
	var genA, genB int64
	genA = 883
	genB = 879
	matching := 0
	matches := &Matches{}

	for pair := 0; pair < 40000000; pair++ {
		genA, genB = generate(genA, genB, matches)

		if genA&0xFFFF == genB&0xFFFF {
			matching++
		}
	}

	for matches.compared < 5000000 {
		genA, genB = generate(genA, genB, matches)
	}

	fmt.Println("Part 1:", matching)
	fmt.Println("Part 2:", matches.matching)
}
