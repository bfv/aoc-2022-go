package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
)

// A, X = rock (lose)     1
// B, Y = paper (draw)    2
// C, Z = scissors (win)  3

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day02"

	p1 := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	p2 := map[string]string{
		"A X": "Z",
		"A Y": "X",
		"A Z": "Y",
		"B X": "X",
		"B Y": "Y",
		"B Z": "Z",
		"C X": "Y",
		"C Y": "Z",
		"C Z": "X",
	}

	strs := aoc.GetStringArray("input.txt")

	for _, s := range strs {
		a += p1[s]
		s2 := string(s[0]) + " " + p2[s] // using a intermediate string here is actually faster than inlining the expression
		b += p1[s2]
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}
