package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	move   byte
	first  int
	second int
}

func dance(line []byte, moves []Move) []byte {
	for _, move := range moves {
		switch move.move {
		case 's':
			line = append(line[len(line)-move.first:], line[:len(line)-move.first]...)
		case 'x':
			tmp := line[move.first]
			line[move.first] = line[move.second]
			line[move.second] = tmp
		case 'p':
			var first, second int
			for index, char := range line {
				if char == byte(move.first) {
					first = index
				} else if char == byte(move.second) {
					second = index
				}
			}
			tmp := line[first]
			line[first] = line[second]
			line[second] = tmp
		}
	}
	return line
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	original := []byte("abcdefghijklmnop")
	line := make([]byte, len(original))
	copy(line, original)
	var moves []Move

	for _, instruction := range strings.Split(scanner.Text(), ",") {
		move := Move{}
		move.move = instruction[0]

		switch move.move {
		case 's':
			move.first, _ = strconv.Atoi(instruction[1:])
		case 'x':
			indices := strings.Split(instruction[1:], "/")
			index1Str, index2Str := indices[0], indices[1]
			move.first, _ = strconv.Atoi(index1Str)
			move.second, _ = strconv.Atoi(index2Str)
		case 'p':
			move.first = int(instruction[1])
			move.second = int(instruction[3])
		}

		moves = append(moves, move)
	}

	fmt.Println("Part 1:", string(dance(line, moves)))

	var cycle int
	for index := 1; true; index++ {
		line = dance(line, moves)
		if bytes.Equal(line, original) {
			cycle = index
			break
		}
	}
	for index := 0; index < 1000000000%cycle; index++ {
		line = dance(line, moves)
	}

	fmt.Println("Part 2:", string(line))
}
