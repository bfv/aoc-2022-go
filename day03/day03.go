package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
)

type rucksack struct {
	C1 string
	C2 string
}

func (r *rucksack) Load(s string) {
	l := len(s)
	r.C1 = s[0 : l/2]
	r.C2 = s[l/2 : l]
}

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day03"

	strs := aoc.GetStringArray("input.txt")

	si := ""
	for i, s := range strs {
		r := rucksack{}
		r.Load(s)
		a += getScore(getIntersection(r.C1, r.C2))
		if i%3 == 0 {
			if si != "" {
				b += getScore(si)
			}
			si = r.C1 + r.C2
		} else {
			si = getIntersection(si, r.C1+r.C2)
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
