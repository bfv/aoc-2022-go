package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
)

var vmax []int = []int{0, 0, 0}

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day01"

	var cal int
	strs := aoc.GetStringArray("inputL.txt")

	for _, s := range strs {

		if s == "" {
			if cal > a {
				a = cal
			}
			checkMax(cal)
			cal = 0
		} else {
			c, _ := strconv.Atoi(s)
			cal += c
		}

	}

	b = calcMax()

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func checkMax(vIn int) {
	if vmax[0] < vIn {
		vmax[0] = vIn
		sort.Ints(vmax)
	}
}

func calcMax() int {
	max := 0
	for _, v := range vmax {
		max += v
	}
	return max
}
