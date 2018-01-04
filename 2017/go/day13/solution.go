package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	depth int
	width int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var firewall []Layer
	severity := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		depth, _ := strconv.Atoi(line[0])
		width, _ := strconv.Atoi(line[1])

		firewall = append(firewall, Layer{
			depth: depth,
			width: width,
		})

		if depth%(width*2-2) == 0 {
			severity += depth * width
		}
	}

	var delay int
	caught := true

	for delay = 0; caught; delay++ {
		caught = false
		for _, layer := range firewall {
			if (layer.depth+delay+1)%(layer.width*2-2) == 0 {
				caught = true
				break
			}
		}
	}

	fmt.Println("Part 1:", severity)
	fmt.Println("Part 2:", delay)
}
