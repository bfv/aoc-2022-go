package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

var treeMap [][]int

func main() {

	start := time.Now()

	day := "day07"
	var a, b int

	input := aoc.GetStringArray("input.txt")
	ints := toArrayOfIntArray(input)

	treeMap = initMap(ints)

	// fmt.Println(calcScenicScore(3, 2, ints))
	// os.Exit(0)
	a = evalTreesA(ints)
	b = evalTreesB(ints)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func toArrayOfIntArray(strs []string) [][]int {
	ints := make([][]int, 0)
	for _, s := range strs {
		ints = append(ints, aoc.GetDigitArray(s))
	}
	return ints
}

func evalTreesA(ints [][]int) int {

	for i := 1; i < len(ints)-1; i++ {
		treeRow := ints[i]
		treeRowVis := evalTreeLine(treeRow)
		applyRow(i, treeRowVis)

		treeCol := getTreeCol(i, ints)
		treeColVis := evalTreeLine(treeCol)
		applyCol(i, treeColVis)
	}

	return countTreeMap()
}

func countTreeMap() int {
	vis := (len(treeMap[0]) - 1) * 4
	for _, row := range treeMap {
		for _, v := range row {
			vis += v
		}
	}
	return vis
}

func getTreeCol(col int, ints [][]int) []int {
	colTrees := []int{}
	for i := 0; i < len(ints); i++ {
		colTrees = append(colTrees, ints[i][col])
	}
	return colTrees
}

func initMap(ints [][]int) [][]int {
	mp := make([][]int, len(ints))
	for i := range mp {
		mp[i] = make([]int, len(ints))
	}
	return mp
}

func evalTreeLine(trees []int) []int {

	maxL := trees[0]
	maxR := trees[len(trees)-1]
	vis := make([]int, len(trees))

	for i := 1; i < (len(trees) - 1); i++ {

		l := trees[i]
		if l > maxL {
			maxL = l
			vis[i] = 1
		}

		r := trees[len(trees)-i-1]
		if r > maxR {
			maxR = r
			vis[len(trees)-i-1] = 1
		}
	}

	return vis
}

func applyRow(row int, trees []int) {
	for i, v := range trees {
		if v > 0 {
			treeMap[row][i] = 1
		}
	}
}

func applyCol(col int, trees []int) {
	for i, v := range trees {
		if v > 0 {
			treeMap[i][col] = 1
		}
	}
}

func evalTreesB(ints [][]int) int {
	m := 0
	for r, row := range ints {
		for c, _ := range row {
			m = lib.Max(m, calcScenicScore(r, c, ints))
		}
	}
	return m
}

func calcScenicScore(row int, col int, ints [][]int) int {
	var c, r, s int
	exit := false

	score := 1
	thh := ints[row][col] // treehouse height

	// up
	exit, s = false, 0
	for r = row - 1; r >= 0 && !exit; r-- {
		s++
		exit = evalScenic(ints[r][col], thh)
	}
	score *= s

	// down
	exit, s = false, 0
	for r = row + 1; r < len(ints) && !exit; r++ {
		s++
		exit = evalScenic(ints[r][col], thh)
	}
	score *= s

	// left
	exit, s = false, 0
	for c = col - 1; c >= 0 && !exit; c-- {
		s++
		exit = evalScenic(ints[row][c], thh)
	}
	score *= s

	// right
	exit, s = false, 0
	for c = col + 1; c < len(ints) && !exit; c++ {
		s++
		exit = evalScenic(ints[row][c], thh)
	}
	score *= s

	return score
}

func evalScenic(th int, thh int) bool {
	return th >= thh
}
