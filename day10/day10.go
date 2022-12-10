package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

var CRT []string

func main() {

	start := time.Now()

	day := "day10"

	commands := aoc.GetStringArray("input.txt")

	x, sumA, ptr, cyclesLeft := 1, 0, 0, 0
	cmd := ""

	CRT = make([]string, 240)

	for cycle := 1; cycle <= 240; cycle++ {

		// start
		if cyclesLeft == 0 {
			cmd = commands[ptr]
			if strings.HasPrefix(cmd, "addx") {
				cyclesLeft = 2
			} else {
				cyclesLeft = 1
			}
			ptr++
		}
		draw(cycle, x)

		// during
		if cycle%40 == 20 {
			strength := x * cycle
			sumA += strength
		}

		//after
		cyclesLeft--
		if cyclesLeft == 0 && strings.HasPrefix(cmd, "addx") {
			x += lib.Atoi(strings.Split(cmd, " ")[1])
		}
	}

	displayCRT()

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, sumA, "EZFPRAKL", elapsed)
}

func displayCRT() {
	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			fmt.Print(CRT[row*40+col])
		}
		fmt.Println()
	}
}

func draw(cycle int, spritePos int) {
	spritePos--
	cycle--
	col := cycle % 40
	if col >= spritePos && col <= spritePos+2 {
		CRT[cycle] = "#"
	} else {
		CRT[cycle] = "."
	}
}
