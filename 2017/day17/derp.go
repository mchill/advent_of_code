package main

import (
	"fmt"
)

func main() {
	const STEPS = 356

	current := 0
	part2 := 0

	for cycle := 1; cycle <= 50000000; cycle++ {
		current = (current + STEPS + 1) % cycle
		if current == 0 {
			part2 = cycle
		}
	}

	fmt.Println("Part 2:", part2)
}
