package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

var queue lib.Stack[rune]

func main() {

	start := time.Now()

	day := "day06"
	var a, b int
	str := aoc.GetStringArray("input.txt")[0]

	queue = lib.Stack[rune]{}
	for i, r := range str {
		addToQueue(r)
		if a == 0 && allUnique(queue, 4) {
			a = i + 1
		}
		if b == 0 && allUnique(queue, 14) {
			b = i + 1
		}
		if a > 0 && b > 0 {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func addToQueue(r rune) {
	queue.AddToBottom(r)
	if queue.Depth() > 14 {
		queue.Pop()
	}
}

func allUnique(q lib.Stack[rune], size int) bool {
	if q.Depth() < size {
		return false
	}
	rs := unique(q.Content()[0:size])
	return len(rs) == size
}

func unique(rs []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range rs {
		if _, v := keys[entry]; !v {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
