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

func parseAction(scanner *bufio.Scanner) Action {
	action := Action{}

	scanner.Scan()
	scanner.Scan()
	action.write = int(strings.Split(scanner.Text(), " ")[8][0] - '0')

	scanner.Scan()
	action.move = 1
	if strings.Split(scanner.Text(), " ")[10] == "left." {
		action.move = -1
	}

	scanner.Scan()
	action.next = strings.Split(scanner.Text(), " ")[8][0]

	return action
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
		scanner.Scan()
		thisState := strings.Split(scanner.Text(), " ")[2][0]

		states[thisState] = State{
			ifFalse: parseAction(scanner),
			ifTrue:  parseAction(scanner),
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
