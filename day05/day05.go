package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
)

func main() {

	start := time.Now()

	var a, b int = 0, 0
	day := "day05"

	strs := aoc.GetStringArray("input.txt")

	for _, s := range strs {
		fmt.Println(s)
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}
