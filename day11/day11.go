package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

type monkey struct {
	nr          int
	items       []int
	operand     string
	rvalue      int
	divider     int
	throwTrue   int
	throwFalse  int
	inspections int
}

func (m *monkey) Receive(newWorry int) {
	m.items = append(m.items, newWorry)
}

var lcm int
var dayPart string

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day01"

	strs := aoc.GetStringArray("input.txt")

	dayPart = "A"
	monkies := parseInput(strs)
	for round := 1; round <= 20; round++ {
		for _, m := range monkies {
			takeTurn(m, monkies)
		}
	}
	a = calcResult(monkies)

	dayPart = "B"
	monkies = parseInput(strs) // fresh set of monkies
	lcm = calcLCM(monkies)
	for round := 1; round <= 10000; round++ {
		for _, m := range monkies {
			takeTurn(m, monkies)
		}
	}
	b = calcResult(monkies)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func displayMonkies(monkies []monkey) {
	for i, m := range monkies {
		fmt.Printf("Monkey %v: (%v) %v\n", i, m.inspections, m.items)
	}
}

func calcResult(monkies []monkey) int {
	results := []int{}
	for _, m := range monkies {
		results = append(results, m.inspections)
	}
	sort.Ints(results)
	return results[len(results)-1] * results[len(results)-2]
}

func parseInput(strs []string) []monkey {
	monkies := []monkey{}

	var m monkey
	for _, s := range strs {
		s = strings.Trim(s, " ")
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "Monkey":
			m = monkey{}
			m.nr = lib.Atoi(strings.ReplaceAll(parts[1], ":", ""))
		case "Starting":
			s = strings.ReplaceAll(s, "Starting items:", "")
			s = strings.ReplaceAll(s, " ", "")
			m.items = lib.ToIntArray(s, ",")
		case "Operation:":
			s = strings.ReplaceAll(s, "Operation: new = old ", "")
			expr := strings.Split(s, " ")
			m.operand = expr[0]
			if expr[1] == "old" {
				m.operand = "^2"
			} else {
				m.rvalue = lib.Atoi(expr[1])
			}
		case "Test:":
			m.divider = lib.Atoi(parts[3])
		case "If":
			if parts[1] == "true:" {
				m.throwTrue = lib.Atoi(parts[5])
			} else {
				m.throwFalse = lib.Atoi(parts[5])
				monkies = append(monkies, m)
			}
		}

	}

	return monkies
}

func takeTurn(m monkey, monkies []monkey) {
	for _, item := range m.items {
		m.inspections++
		newWorry := calcWorry(m, item)
		if test(m, newWorry) {
			monkies[m.throwTrue].Receive(newWorry)
		} else {
			monkies[m.throwFalse].Receive(newWorry)
		}
	}
	m.items = []int{}
	monkies[m.nr] = m
}

func calcWorry(m monkey, old int) int {
	var newWorry int

	switch m.operand {
	case "+":
		newWorry = old + m.rvalue
	case "*":
		newWorry = old * m.rvalue
	case "^2":
		newWorry = old * old
	}

	if dayPart == "A" {
		newWorry = int(math.Floor(float64(newWorry) / 3))
	} else {
		newWorry = newWorry % lcm
	}

	return newWorry
}

func test(m monkey, worry int) bool {
	return worry%m.divider == 0
}

func calcLCM(monkies []monkey) int {
	res := 1
	for _, m := range monkies {
		res *= m.divider
	}
	return res
}
