package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	write int
	move  int
	next  byte
}

type State struct {
	ifFalse Action
	ifTrue  Action
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	states := make(map[byte]State)

	scanner.Scan()
	state := strings.Split(scanner.Text(), " ")[3][0]

	scanner.Scan()
	steps, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[5])

	for scanner.Scan() {
		ifFalse := Action{}
		ifTrue := Action{}

		scanner.Scan()
		thisState := strings.Split(scanner.Text(), " ")[2][0]

		scanner.Scan()
		scanner.Scan()
		ifFalse.write = int(strings.Split(scanner.Text(), " ")[8][0] - '0')

		scanner.Scan()
		ifFalse.move = 1
		if strings.Split(scanner.Text(), " ")[10] == "left." {
			ifFalse.move = -1
		}

		scanner.Scan()
		ifFalse.next = strings.Split(scanner.Text(), " ")[8][0]

		scanner.Scan()
		scanner.Scan()
		ifTrue.write = int(strings.Split(scanner.Text(), " ")[8][0] - '0')

		scanner.Scan()
		ifTrue.move = 1
		if strings.Split(scanner.Text(), " ")[10] == "left." {
			ifTrue.move = -1
		}

		scanner.Scan()
		ifTrue.next = strings.Split(scanner.Text(), " ")[8][0]

		states[thisState] = State{
			ifFalse: ifFalse,
			ifTrue:  ifTrue,
		}
	}

	tape := make(map[int]bool)
	current := 0

	for step := 0; step < steps; step++ {
		action := states[state].ifFalse
		if tape[current] {
			action = states[state].ifTrue
		}

		tape[current] = true
		if action.write == 0 {
			delete(tape, current)
		}
		current += action.move
		state = action.next
	}

	fmt.Println("Part 1:", len(tape))
}
