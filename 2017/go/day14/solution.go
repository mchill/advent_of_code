package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(list []int, current int, length int) {
	max := len(list)
	for i, j := current, current+length-1; i < j; i, j = i+1, j-1 {
		list[i%max], list[j%max] = list[j%max], list[i%max]
	}
}

func knotHash(input string) string {
	var list []int
	for index := 0; index < 256; index++ {
		list = append(list, index)
	}

	var lengths []int
	for _, char := range input {
		lengths = append(lengths, int(char))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	current := 0
	skip := 0
	for index := 0; index < 64; index++ {
		for _, length := range lengths {
			reverse(list, current, length)
			current += length + skip
			skip++
		}
	}

	result := ""
	for i := 0; i < len(list); i += 16 {
		compressed := list[i]
		for j := i + 1; j < i+16; j++ {
			compressed ^= list[j]
		}

		converted := strconv.FormatInt(int64(compressed), 2)
		for j := len(converted); j < 8; j++ {
			converted = "0" + converted
		}
		result += converted
	}

	return result
}

func removeRegion(square string, disk map[string][]string) {
	reachableSquares := disk[square]
	delete(disk, square)

	for _, reachable := range reachableSquares {
		removeRegion(reachable, disk)
	}
}

func main() {
	input := "ljoxqyyw-"
	used := 0
	regions := 0
	disk := make(map[string][]string)

	for row := 0; row < 128; row++ {
		hash := knotHash(input + strconv.Itoa(row))
		used += strings.Count(hash, "1")

		for index, char := range hash {
			if char == '0' {
				continue
			}

			key := string(row) + string(index)

			disk[key] = append(disk[key], string(row-1)+string(index))
			disk[key] = append(disk[key], string(row+1)+string(index))
			disk[key] = append(disk[key], string(row)+string(index-1))
			disk[key] = append(disk[key], string(row)+string(index+1))
		}
	}

	for len(disk) > 0 {
		var square string
		for square, _ = range disk {
			break
		}
		regions++
		removeRegion(square, disk)
	}

	fmt.Println("Part 1:", used)
	fmt.Println("Part 2:", regions)
}
