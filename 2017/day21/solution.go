package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func rotate(input string) string {
	rows := strings.Split(input, "/")
	newRows := make([]string, len(rows))
	for i := 0; i < len(rows[0]); i++ {
		for j := len(rows[0]) - 1; j >= 0; j-- {
			newRows[i] += string(rows[j][i])
		}
	}
	return strings.Join(newRows, "/")
}

func flip(input string) string {
	rows := strings.Split(input, "/")
	rows[0], rows[len(rows)-1] = rows[len(rows)-1], rows[0]
	return strings.Join(rows, "/")
}

func split(input string) [][]string {
	rows := strings.Split(input, "/")
	size := len(rows[0])
	width := 3
	if size%2 == 0 {
		width = 2
	}

	count := size / width
	inputs := make([][]string, count*count)

	for split := 0; split < count*count; split++ {
		row := split / count
		col := split % count
		for i := row * width; i < (row+1)*width; i++ {
			inputs[split] = append(inputs[split], rows[i][col*width:(col+1)*width])
		}
	}

	return inputs
}

func combine(input [][]string) string {
	count := int(math.Sqrt(float64(len(input))))
	width := len(input[0])

	for section := 0; section < count; section++ {
		for offset := 1; offset < count; offset++ {
			input[section] = append(input[section], input[section+count*offset]...)
		}
	}

	for row := 0; row < count*width; row++ {
		for section := 1; section < count; section++ {
			input[0][row] += input[section][row]
		}
	}

	return strings.Join(input[0], "/")
}

func transform(inputs [][]string, rules map[string]string) {
	for index, section := range inputs {
		output := rules[strings.Join(section, "/")]
		inputs[index] = strings.Split(output, "/")
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	rules := make(map[string]string)

	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), " => ")
		input, output := rule[0], rule[1]
		for i := 0; i < 2; i++ {
			for j := 0; j < 4; j++ {
				rules[input] = output
				input = rotate(input)
			}
			input = flip(input)
		}
	}

	input := ".#./..#/###"

	for i := 0; i < 5; i++ {
		inputs := split(input)
		transform(inputs, rules)
		input = combine(inputs)
	}

	fmt.Println("Part 1:", strings.Count(input, "#"))

	for i := 0; i < 13; i++ {
		inputs := split(input)
		transform(inputs, rules)
		input = combine(inputs)
	}

	fmt.Println("Part 2:", strings.Count(input, "#"))
}
