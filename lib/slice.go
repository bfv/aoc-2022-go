package lib

// no indexOf function for slices. we wait for generics :-)
func IndexOfString(toFind string, list []string) int {
	for k, v := range list {
		if toFind == v {
			return k
		}
	}
	return -1 //not found.
}
