package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoclib"
)

var CRT []string

func main() {

	start := time.Now()
	day := "day10"

	commands := aoclib.GetStringArray("input.txt")

	x, sumA, ptr, cyclesLeft := 1, 0, 0, 0
	var cmd []string

	CRT = make([]string, 240)

	for cycle := 1; cycle <= 240; cycle++ {

		// start
		if cyclesLeft == 0 {
			cmd = strings.Split(commands[ptr], " ")
			switch cmd[0] {
			case "noop":
				cyclesLeft = 1
			case "addx":
				cyclesLeft = 2
			}
			ptr++
		}

		// during
		if cycle%40 == 20 {
			sumA += x * cycle
		}
		draw(cycle, x)

		//after
		cyclesLeft--
		if cyclesLeft == 0 {
			switch cmd[0] {
			case "addx":
				x += aoclib.Atoi(cmd[1])
			}
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
	cycle-- // from 1-based to 0-based
	spritePos--
	col := cycle % 40
	if col >= spritePos && col <= spritePos+2 {
		CRT[cycle] = "#"
	} else {
		CRT[cycle] = "."
	}
}
