package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(list []int, current int, length int) {
	max := len(list)
	for i, j := current, current+length-1; i < j; i, j = i+1, j-1 {
		list[i%max], list[j%max] = list[j%max], list[i%max]
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, ",")

	var list1, list2 []int
	current := 0
	skip := 0

	for index := 0; index < 256; index++ {
		list1 = append(list1, index)
		list2 = append(list2, index)
	}

	for _, word := range words {
		length, _ := strconv.Atoi(word)
		reverse(list1, current, length)
		current += length + skip
		skip++
	}

	current = 0
	skip = 0

	var lengths []int
	for _, char := range line {
		lengths = append(lengths, int(char))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	for index := 0; index < 64; index++ {
		for _, length := range lengths {
			reverse(list2, current, length)
			current += length + skip
			skip++
		}
	}

	hex := ""
	for i := 0; i < len(list2); i += 16 {
		compressed := list2[i]
		for j := i + 1; j < i+16; j++ {
			compressed ^= list2[j]
		}

		converted := strconv.FormatInt(int64(compressed), 16)
		if len(converted) == 1 {
			hex += "0"
		}
		hex += converted
	}

	fmt.Println("Part 1:", list1[0]*list1[1])
	fmt.Println("Part 2:", hex)
}
