package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoclib"
)

var signal string

func main() {

	start := time.Now()

	day := "day06"
	var a, b int
	signal = aoclib.GetStringArray("input.txt")[0]

	for pos := range signal {
		if a == 0 && allUnique(pos, 4) {
			a = pos
		}
		if b == 0 && allUnique(pos, 14) {
			b = pos
		}
		if a > 0 && b > 0 {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func allUnique(pos int, size int) bool {
	if pos < size {
		return false
	}
	rs := unique(signal[pos-size : pos])
	return len(rs) == size
}

func unique(rs string) []rune {
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
