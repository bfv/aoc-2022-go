package aoc

import (
	"io/ioutil"
	"math"
	"runtime"
	"strconv"
	"strings"

	"github.com/bfv/aoc2022-go/lib"
)

var LineEnd string

func init() {
	if runtime.GOOS == "windows" {
		LineEnd = "\r\n"
	} else {
		LineEnd = "\n"
	}
}

func GetStringArray(filename string) []string {

	buf, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(buf), LineEnd)

	return input
}

func GetIntArray(filename string) []int {
	ints := []int{}
	lines := GetStringArray(filename)
	for _, line := range lines {
		x, _ := strconv.Atoi(line)
		ints = append(ints, x)
	}
	return ints
}

func GetIntArrayFromBinary(filename string) []int {
	ints := []int{}
	lines := GetStringArray(filename)
	for _, line := range lines {
		x, _ := strconv.ParseInt(line, 2, 64)
		ints = append(ints, int(x))
	}
	return ints
}

func GetIntArrayFromStringArray(inputStrings []string) []int {
	array := []int{}
	for _, f := range inputStrings {
		array = append(array, lib.Atoi(f))
	}
	return array
}

func GetIntArrayFromString(inputString string) []int {
	inputStrings := strings.Split(inputString, "")
	return GetIntArrayFromStringArray(inputStrings)
}

func GetInputLineLength(filename string) int {
	lines := GetStringArray(filename)
	return len(lines[0])
}

func GetNumberArray(line string) []int {
	nmbrs := strings.Split(line, ",")
	numbers := make([]int, len(nmbrs))
	for i, v := range nmbrs {
		x, _ := strconv.Atoi(v)
		numbers[i] = x
	}
	return numbers
}

func GetDigitArray(line string) []int {
	nmbrs := strings.Split(line, "")
	numbers := make([]int, len(nmbrs))
	for i, v := range nmbrs {
		x, _ := strconv.Atoi(v)
		numbers[i] = x
	}
	return numbers
}

func GetHighestIntFromArray(ints []int) int {
	h := math.MinInt
	for _, x := range ints {
		if x > h {
			h = x
		}
	}
	return h
}
