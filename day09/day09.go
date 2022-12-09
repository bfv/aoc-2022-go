package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

type position struct {
	x int
	y int
}

var head position
var tail position
var positions map[string]int
var rope []position // rope[0] is head, rope[9] is tail
var part string

func main() {

	start := time.Now()

	day := "day09"
	var a, b int

	head = position{0, 0}
	tail = position{0, 0}
	positions = map[string]int{}
	registerTail(tail)

	input := aoc.GetStringArray("input.txt")

	// a
	part = "A"
	for _, m := range input {
		head = moveHeadA(m)
	}
	a = countPositions()

	// b
	part = "B"
	positions = map[string]int{}
	rope = make([]position, 10)

	for i, m := range input {
		moveHeadB(m, i)
	}

	b = countPositions()

	elapsed := time.Since(start)
	fmt.Printf("\n\n%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func moveHeadA(cmd string) position {
	parts := strings.Split(cmd, " ")
	dir := parts[0]
	amount := lib.Atoi(parts[1])
	for i := 0; i < amount; i++ {
		head = moveKnot(head, dir)
		tail = evalTail(dir, head, tail)
	}
	return head
}

func moveHeadB(cmd string, nrCmd int) {

	parts := strings.Split(cmd, " ")
	dir := parts[0]
	amount := lib.Atoi(parts[1])

	for i := 0; i < amount; i++ {

		rope[0] = moveKnot(rope[0], dir) // move the head

		for j := 1; j < 10; j++ { // and the rest will follow

			headB := rope[j-1]
			tailB := rope[j]

			dx, dy := headB.x-tailB.x, headB.y-tailB.y
			adx, ady := lib.Abs(dx), lib.Abs(dy)

			if adx == 2 && dy == 0 {
				rope[j].x += dx / 2
			}

			if ady == 2 && dx == 0 {
				rope[j].y += dy / 2
			}

			if (adx == 1 && ady == 2) ||
				(adx == 2 && ady == 1) ||
				(adx == 2 && ady == 2) {

				if dx > 0 {
					rope[j].x++
				} else {
					rope[j].x--
				}

				if dy > 0 {
					rope[j].y++
				} else {
					rope[j].y--
				}
			}
		}
		registerTail(rope[9])
	}
}

func moveKnot(knot position, dir string) position {
	switch dir {
	case "U":
		knot.y++
	case "D":
		knot.y--
	case "L":
		knot.x--
	case "R":
		knot.x++
	}
	return position{knot.x, knot.y}
}

func evalTail(dir string, head position, tail position) position {

	dist := calcDistance(head, tail)

	if dist == 2.0 {
		switch dir {
		case "U":
			tail = position{tail.x, tail.y + 1}
		case "D":
			tail = position{tail.x, tail.y - 1}
		case "L":
			tail = position{tail.x - 1, tail.y}
		case "R":
			tail = position{tail.x + 1, tail.y}
		}
	} else if dist > 2.0 {
		switch dir {
		case "U":
			tail = position{head.x, head.y - 1}
		case "D":
			tail = position{head.x, head.y + 1}
		case "L":
			tail = position{head.x + 1, head.y}
		case "R":
			tail = position{head.x - 1, head.y}
		}
	}

	if dist >= 2.0 && part == "A" {
		registerTail(tail)
	}

	return tail
}

func calcDistance(head position, tail position) float64 {
	return math.Sqrt(math.Pow(float64(head.x)-float64(tail.x), 2) + math.Pow(float64(head.y)-float64(tail.y), 2))
}

func registerTail(tail position) {
	c := fmt.Sprintf("%v,%v", tail.x, tail.y)
	positions[c] += 1
}

func countPositions() int {
	return len(positions)
}

func displayBoard(sizeX int, sizeY int) {

	for y := sizeY; y >= 0; y-- {

		for x := 0; x < sizeX; x++ {
			printed := false
			for i := 0; i < 10 && !printed; i++ {
				if rope[i].x == x && rope[i].y == y {
					if i == 0 {
						fmt.Printf("H")
					} else {
						fmt.Printf("%v", i)
					}
					printed = true
				}
			}
			if !printed {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
