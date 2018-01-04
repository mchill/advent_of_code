package main

import (
	"fmt"
	"math"
)

func abs(number int) int {
	if number < 0 {
		number = -number
	}
	return number
}

func exp(number int, exponent int) int {
	result := number
	for index := 1; index < exponent; index++ {
		result *= number
	}
	return result
}

func part1(input int) {
	width := int(math.Ceil(math.Sqrt(float64(input))))
	if width%2 == 0 {
		width++
	}

	part1 := abs((input-exp(width-2, 2))%(width-1)-width/2) + width/2
	fmt.Println("Part 1:", part1)
}

func part2(input int) {
	DIRECTIONS := [][]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}

	values := map[string]int{"0,0": 1}
	direction := 3
	distance := 2
	counter := distance
	xpos := 0
	ypos := 0

	for {
		if counter % (distance/2) == 0 {
			direction = (direction + 1) % len(DIRECTIONS)
		}
		if counter == 0 {
			distance += 2
			counter = distance
		}
		counter--

		xpos += DIRECTIONS[direction][0]
		ypos += DIRECTIONS[direction][1]

		value := values[fmt.Sprintf("%d%s%d", xpos, ",", ypos)] +
			values[fmt.Sprintf("%d%s%d", xpos+1, ",", ypos)] +
			values[fmt.Sprintf("%d%s%d", xpos+1, ",", ypos-1)] +
			values[fmt.Sprintf("%d%s%d", xpos, ",", ypos-1)] +
			values[fmt.Sprintf("%d%s%d", xpos-1, ",", ypos-1)] +
			values[fmt.Sprintf("%d%s%d", xpos-1, ",", ypos)] +
			values[fmt.Sprintf("%d%s%d", xpos-1, ",", ypos+1)] +
			values[fmt.Sprintf("%d%s%d", xpos, ",", ypos+1)] +
			values[fmt.Sprintf("%d%s%d", xpos+1, ",", ypos+1)]

		if value > input {
			fmt.Println("Part 2:", value)
			return
		}

		values[fmt.Sprintf("%d%s%d", xpos, ",", ypos)] = value
	}
}

func main() {
	input := 347991
	part1(input)
	part2(input)
}
