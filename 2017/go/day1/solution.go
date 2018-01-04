package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading input.txt")
		panic(err)
	}

	length := len(data) - 1
	data = data[:length]
	part1 := 0
	part2 := 0

	for index := 0; index < length; index++ {
		if data[index] == data[(index+1)%length] {
			part1 += int(data[index]) - '0'
		}
		if data[index] == data[(index+length/2)%length] {
			part2 += int(data[index]) - '0'
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

