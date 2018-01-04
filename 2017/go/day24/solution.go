package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Component struct {
	port1 int
	port2 int
	used  bool
}

func build(lastPort int, components map[int][]*Component, longest bool) (int, int) {
	length := 0
	strength := 0

	for _, component := range components[lastPort] {
		if component.used {
			continue
		}

		nextPort := component.port2
		if lastPort != component.port1 {
			nextPort = component.port1
		}

		component.used = true
		addedLength, addedStrength := build(nextPort, components, longest)
		component.used = false

		addedLength += 1
		addedStrength += component.port1 + component.port2

		if (!longest && addedStrength > strength) ||
			(longest && addedLength > length) ||
			(addedLength == length && addedStrength > strength) {
			length = addedLength
			strength = addedStrength
		}
	}

	return length, strength
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	components := make(map[int][]*Component)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "/")
		port1, _ := strconv.Atoi(line[0])
		port2, _ := strconv.Atoi(line[1])

		component := &Component{
			port1: port1,
			port2: port2,
		}

		components[port1] = append(components[port1], component)
		components[port2] = append(components[port2], component)
	}

	_, strongest := build(0, components, false)
	fmt.Println("Part 1:", strongest)

	_, strongest = build(0, components, true)
	fmt.Println("Part 2:", strongest)
}
