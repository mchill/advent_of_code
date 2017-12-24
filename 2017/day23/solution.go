package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	inst   string
	first  interface{}
	second interface{}
}

func value(val interface{}, registers map[byte]int) int {
	switch val.(type) {
	case byte:
		return registers[val.(byte)]
	case int:
		return val.(int)
	}
	return 0
}

func parseFile() []Instruction {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var instructions []Instruction

	for scanner.Scan() {
		line := strings.Split(scanner.Text()+" |", " ")

		var first, second interface{}
		first, err = strconv.Atoi(line[1])
		if err != nil {
			first = line[1][0]
		}
		second, err = strconv.Atoi(line[2])
		if err != nil {
			second = line[2][0]
		}

		instruction := Instruction{
			inst:   line[0],
			first:  first,
			second: second,
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func part1() {
	instructions := parseFile()
	registers := make(map[byte]int)
	current := 0
	count := 0

	for current < len(instructions) {
		inst := instructions[current]
		current++

		switch inst.inst {
		case "set":
			registers[inst.first.(byte)] = value(inst.second, registers)
		case "sub":
			registers[inst.first.(byte)] -= value(inst.second, registers)
		case "mul":
			registers[inst.first.(byte)] *= value(inst.second, registers)
			count++
		case "jnz":
			if value(inst.first, registers) != 0 {
				current += value(inst.second, registers) - 1
			}
		}
	}

	fmt.Println("Part 1:", count)
}

func part2() {
	h := 0
	for b := 93*100 + 100000; b < 93*100 + 100000 + 17000; b += 17 {
		if !big.NewInt(int64(b)).ProbablyPrime(1) {
			h++
		}
	}
	fmt.Println("Part 2:", h+1)
}

func main() {
	part1()
	part2()
}
