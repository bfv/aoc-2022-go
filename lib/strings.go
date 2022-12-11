package lib

import "strings"

func GetUniqueRunes(s string) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range s {
		if _, v := keys[entry]; !v {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ToIntArray(s string, sep string) []int {
	ints := []int{}
	intsStrs := strings.Split(s, sep)
	for _, in := range intsStrs {
		ints = append(ints, Atoi(in))
	}
	return ints
}
