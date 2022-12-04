package lib

import (
	"strconv"
	"strings"
)

// If you're sure it's an int
func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.Trim(s, " "))
	return i
}
