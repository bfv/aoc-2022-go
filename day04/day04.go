package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
)

type assignment struct {
	Low  int
	High int
}

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day04"

	strs := aoc.GetStringArray("input.txt")

	for _, s := range strs {
		pair := strings.Split(s, ",")
		e1, e2 := expand(pair[0]), expand(pair[1])
		if isContained(e1, e2) {
			a++
		}
		if isOverlapping(e1, e2) {
			b++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func expand(s string) assignment {
	vs := strings.Split(s, "-")
	i1, _ := strconv.Atoi(vs[0])
	i2, _ := strconv.Atoi(vs[1])
	return assignment{i1, i2}
}

func isContained(a1, a2 assignment) bool {
	return a1.Low <= a2.Low && a1.High >= a2.High || a2.Low <= a1.Low && a2.High >= a1.High
}

func isOverlapping(a1, a2 assignment) bool {
	return a1.Low >= a2.Low && a1.Low <= a2.High ||
		a1.High >= a2.Low && a1.High <= a2.High ||
		a1.Low <= a2.Low && a1.High >= a2.High
}
