package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	supported       []*Program
	supportedNames  []string
	weight          int
	supportedWeight int
}

func findRoot(
	programs map[string]*Program,
	nodes map[string]bool,
	leaves []string) (string, *Program) {
	var name string

	for _, leaf := range leaves {
		delete(nodes, leaf)
	}
	for name, _ = range nodes {
		break
	}

	return name, programs[name]
}

func calculateWeight(program *Program) int {
	weight := 0

	for _, supported := range program.supported {
		weight += calculateWeight(supported)
	}

	program.supportedWeight = weight
	return program.weight + program.supportedWeight
}

func balance(program *Program, difference int) {
	programs := make([][]*Program, 2)
	var weight int

	for _, supported := range program.supported {
		totalWeight := supported.weight + supported.supportedWeight

		if weight == 0 {
			weight = totalWeight
		}

		if weight == totalWeight {
			programs[0] = append(programs[0], supported)
		} else {
			programs[1] = append(programs[1], supported)
		}
	}

	if len(programs[1]) == 0 {
		fmt.Println("Part 2:", program.weight + difference)
		return
	}

	firstWeight := programs[0][0].weight + programs[0][0].supportedWeight
	secondWeight := programs[1][0].weight + programs[1][0].supportedWeight
	difference = firstWeight - secondWeight

	if len(programs[0]) == 1 {
		balance(programs[0][0], -difference)
	}
	if len(programs[1]) == 1 {
		balance(programs[1][0], difference)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	programs := make(map[string]*Program)
	nodes := make(map[string]bool)
	var leaves []string

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		name := parts[0]
		weight, _ := strconv.Atoi(strings.Trim(parts[1], "()"))

		program := Program{
			weight: weight,
		}
		programs[name] = &program
		nodes[name] = true

		if len(parts) == 2 {
			continue
		}

		for _, part := range parts[3:] {
			supportedProgram := strings.TrimRight(part, ",")
			leaves = append(leaves, supportedProgram)
			program.supportedNames = append(program.supportedNames, supportedProgram)
		}
	}

	for _, program := range programs {
		for _, supported := range program.supportedNames {
			program.supported = append(program.supported, programs[supported])
		}
	}

	part1, root := findRoot(programs, nodes, leaves)
	calculateWeight(root)

	fmt.Println("Part 1:", part1)
	balance(root, 0)
}
