package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
	z int
}

type Particle struct {
	position     *Coordinate
	velocity     *Coordinate
	acceleration *Coordinate
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func constructCoordinate(input string) *Coordinate {
	parts := strings.Split(input, ",")
	xStr, yStr, zStr := parts[0], parts[1], parts[2]
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	z, _ := strconv.Atoi(zStr)
	return &Coordinate{
		x: x,
		y: y,
		z: z,
	}
}

func getManhattan(coordinate *Coordinate) int {
	return abs(coordinate.x) + abs(coordinate.y) + abs(coordinate.z)
}

func tick(particle *Particle) {
	p := particle.position
	v := particle.velocity
	a := particle.acceleration

	v.x, v.y, v.z = v.x+a.x, v.y+a.y, v.z+a.z
	p.x, p.y, p.z = p.x+v.x, p.y+v.y, p.z+v.z
}

func part1(particles map[int]*Particle) int {
	smallestAcceleration := []int{0}
	slowest := getManhattan(particles[0].acceleration)

	for id, particle := range particles {
		acceleration := getManhattan(particle.acceleration)
		if acceleration < slowest {
			smallestAcceleration = []int{id}
			slowest = acceleration
		} else if acceleration == slowest {
			smallestAcceleration = append(smallestAcceleration, id)
		}
	}

	smallestVelocity := smallestAcceleration[0]
	slowest = getManhattan(particles[smallestAcceleration[0]].velocity)

	for _, id := range smallestAcceleration {
		velocity := getManhattan(particles[id].velocity)
		if velocity < slowest {
			smallestVelocity = id
			slowest = velocity
		}
	}

	return smallestVelocity
}

func part2(particles map[int]*Particle) int {
	collisions := make(map[string][]int)

	for i := 0; i < 100; i++ {
		for id, particle := range particles {
			position := particle.position
			index := fmt.Sprintf("%d,%d,%d", position.x, position.y, position.z)
			collisions[index] = append(collisions[index], id)
			tick(particle)
		}

		for _, collision := range collisions {
			if len(collision) == 1 {
				continue
			}
			for _, id := range collision {
				delete(particles, id)
			}
		}
	}

	return len(particles)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	particles := make(map[int]*Particle)
	id := -1

	for scanner.Scan() {
		id++
		line := strings.Replace(scanner.Text(), ">", "", -1)
		line = strings.Replace(line, "<", ", ", -1)
		sets := strings.Split(line, ", ")

		particle := &Particle{
			position:     constructCoordinate(sets[1]),
			velocity:     constructCoordinate(sets[3]),
			acceleration: constructCoordinate(sets[5]),
		}
		particles[id] = particle
	}

	fmt.Println("Part 1:", part1(particles))
	fmt.Println("Part 2:", part2(particles))
}
