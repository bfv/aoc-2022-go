package lib

import (
	"strconv"
	"strings"
)

// If you're sure it's an int
func Atoi(s string) int {
	s = strings.Trim(s, " ")
	i, _ := strconv.Atoi(s)
	return i
}
