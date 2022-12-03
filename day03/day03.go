package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
)

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day03"

	strs := aoc.GetStringArray("input.txt")

	si := ""
	for i, s := range strs {
		a += getScore(getIntersection(s[0:(len(s)/2)], s[(len(s)/2):]))
		if i%3 == 0 {
			if si != "" {
				b += getScore(si)
			}
			si = s
		} else {
			si = getIntersection(si, s)
		}
	}
	b += getScore(si)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func getScore(s string) int {
	var score int
	r := rune(s[0])
	if r >= rune('a') && r <= rune('z') {
		score = int(r) - int(rune('a')) + 1
	} else {
		score = int(r) - int(rune('A')) + 27
	}
	return score
}

func getIntersection(s1, s2 string) string {
	si := ""
	for _, r1 := range s1 {
		if strings.Contains(s2, string(r1)) && !strings.Contains(si, string(r1)) {
			si += string(r1)
		}
	}
	return si
}
