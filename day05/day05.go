package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

var stacksA []lib.Stack[string]
var stacksB []lib.Stack[string]

func main() {

	start := time.Now()

	day := "day05"
	var a, b string
	strs := aoc.GetStringArray("input.txt")

	stacksA = make([]lib.Stack[string], (len(strs[0])+1)/4)
	stacksB = make([]lib.Stack[string], (len(strs[0])+1)/4)

	for _, s := range strs {
		processRow(s)
	}

	a, b = getTops(stacksA), getTops(stacksB)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func processRow(s string) {
	if s != "" {
		c := strings.Trim(s, " ")[0:1]
		switch c {
		case "[":
			processCrateRow(s)
		case "m":
			processMove(s)
		}
	}
}

func processCrateRow(s string) {
	stackNr := 0
	for i, c := range s {
		if i%4 == 1 {
			if string(c) != " " {
				stacksA[stackNr].AddToBottom(string(c))
				stacksB[stackNr].AddToBottom(string(c))
			}
			stackNr++
		}
	}
}

func processMove(s string) {

	tmp := lib.Stack[string]{}
	cmd := strings.Split(s, " ")
	cnt := lib.Atoi(cmd[1])
	from, to := lib.Atoi(cmd[3])-1, lib.Atoi(cmd[5])-1

	for i := 0; i < cnt; i++ {
		stacksA[to].Push(stacksA[from].Pop())
		tmp.Push(stacksB[from].Pop())
	}

	for i := 0; i < cnt; i++ {
		stacksB[to].Push(tmp.Pop())
	}
}

func getTops(stack []lib.Stack[string]) string {
	s := ""
	for i := 0; i < len(stacksA); i++ {
		s += fmt.Sprintf("%v", stack[i].Pop())
	}
	return s
}
