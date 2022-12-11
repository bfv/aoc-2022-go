package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoclib"
)

var stacksA []aoclib.Stack[string]
var stacksB []aoclib.Stack[string]

func main() {

	start := time.Now()

	day := "day05"
	var a, b string
	strs := aoclib.GetStringArray("input.txt")

	stacksA = make([]aoclib.Stack[string], (len(strs[0])+1)/4)
	stacksB = make([]aoclib.Stack[string], (len(strs[0])+1)/4)

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

	tmp := aoclib.Stack[string]{}
	cmd := strings.Split(s, " ")
	cnt := aoclib.Atoi(cmd[1])
	from, to := aoclib.Atoi(cmd[3])-1, aoclib.Atoi(cmd[5])-1

	for i := 0; i < cnt; i++ {
		stacksA[to].Push(stacksA[from].Pop())
		tmp.Push(stacksB[from].Pop())
	}

	for i := 0; i < cnt; i++ {
		stacksB[to].Push(tmp.Pop())
	}
}

func getTops(stack []aoclib.Stack[string]) string {
	s := ""
	for i := 0; i < len(stacksA); i++ {
		s += fmt.Sprintf("%v", stack[i].Pop())
	}
	return s
}
