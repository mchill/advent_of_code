package main

import (
	"container/ring"
	"fmt"
)

func main() {
	const STEPS = 356

	spinlock := ring.New(1)
	spinlock.Value = 0

	for cycle := 1; cycle <= 2017; cycle++ {
		for step := 0; step < STEPS%cycle; step++ {
			spinlock = spinlock.Next()
		}
		spinlock = spinlock.Link(ring.New(1)).Prev()
		spinlock.Value = cycle
	}

	current := 0
	part2 := 0

	for cycle := 1; cycle <= 50000000; cycle++ {
		current = (current + STEPS + 1) % cycle
		if current == 0 {
			part2 = cycle
		}
	}

	fmt.Println("Part 1:", spinlock.Next().Value)
	fmt.Println("Part 2:", part2)
}
