package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func part1(instructions []Instruction) {
	registers := make(map[byte]int)
	current := 0
	var last int

	for instructions[current].inst != "rcv" ||
		value(instructions[current].first, registers) == 0 {
		inst := instructions[current]

		switch inst.inst {
		case "snd":
			last = value(inst.first, registers)
		case "set":
			registers[inst.first.(byte)] = value(inst.second, registers)
		case "add":
			registers[inst.first.(byte)] += value(inst.second, registers)
		case "mul":
			registers[inst.first.(byte)] *= value(inst.second, registers)
		case "mod":
			registers[inst.first.(byte)] %= value(inst.second, registers)
		case "jgz":
			if value(inst.first, registers) > 0 {
				current += value(inst.second, registers) - 1
			}
		}

		current = (current + 1) % len(instructions)
	}

	fmt.Println("Part 1:", last)
}

func part2(instructions []Instruction, ch1 chan int, ch2 chan int,
	quit chan bool, wait *sync.WaitGroup, id int) {
	defer wait.Done()

	registers := make(map[byte]int)
	registers['p'] = id
	current := 0
	sent := 0

	for true {
		inst := instructions[current]

		switch inst.inst {
		case "snd":
			ch1 <- value(inst.first, registers)
			sent++
		case "rcv":
			if len(ch1) == 0 && len(ch2) == 0 {
				quit <- true
				quit <- true
			}
			select {
			case registers[inst.first.(byte)] = <-ch2:
			case <-quit:
				if id == 1 {
					fmt.Println("Part 2:", sent)
				}
				return
			}
		case "set":
			registers[inst.first.(byte)] = value(inst.second, registers)
		case "add":
			registers[inst.first.(byte)] += value(inst.second, registers)
		case "mul":
			registers[inst.first.(byte)] *= value(inst.second, registers)
		case "mod":
			registers[inst.first.(byte)] %= value(inst.second, registers)
		case "jgz":
			if value(inst.first, registers) > 0 {
				current += value(inst.second, registers) - 1
			}
		}

		current = (current + 1) % len(instructions)
	}
}

func main() {
	instructions := parseFile()
	part1(instructions)

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	quit := make(chan bool, 2)

	var wait sync.WaitGroup
	wait.Add(2)
	go part2(instructions, ch1, ch2, quit, &wait, 0)
	go part2(instructions, ch2, ch1, quit, &wait, 1)
	wait.Wait()
}
