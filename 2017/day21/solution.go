package main

import (
	"bufio"
	"fmt"
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

func split(input string) ([][][]string, int) {
	rows := strings.Split(input, "/")
	size := len(rows[0])
	width := 3
	if size%2 == 0 {
		width = 2
	}

	count := size/width
	inputs := make([][][]string, count)
	for col := 0; col < count; col++ {
		inputs[col] = make([][]string, count)
	}

	for split := 0; split < count*count; split++ {
		row := split / count
		col := split % count
		for i := row*width; i < (row+1)*width; i++ {
			inputs[row][col] = append(inputs[row][col], rows[i][col*width:(col+1)*width])
		}
	}

	return inputs, count
}

func combine(inputs [][][]string, count int) string {
	width := len(inputs[0][0])
	var output string

	for col := 0; col < count; col++ {
		for row := 0; row < width; row++ {
			for layer := 0; layer < count; layer++ {
				output += inputs[col][layer][row]
			}
			output += "/"
		}
	}

	return output[:len(output)-1]
}

func transform(inputs [][][]string, rules map[string]string) {
	for rowIndex, row := range inputs {
		for colIndex, col := range row {
			output := rules[strings.Join(col, "/")]
			inputs[rowIndex][colIndex] = strings.Split(output, "/")
		}
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
		inputs, count := split(input)
		transform(inputs, rules)
		input = combine(inputs, count)
	}

	fmt.Println("Part 1:", strings.Count(input, "#"))

	for i := 0; i < 13; i++ {
                inputs, count := split(input)
                transform(inputs, rules)
                input = combine(inputs, count)
		fmt.Println(i+5)
        }

        fmt.Println("Part 2:", strings.Count(input, "#"))
}
