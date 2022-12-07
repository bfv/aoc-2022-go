package lib

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
