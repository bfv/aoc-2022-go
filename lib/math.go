package lib

// since the math package only does float32/64

// Returns the absolute value of an int
func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

// having multiple return paths is ugly...

// Returns the maximum of two ints
func Max(a, b int) int {
	if b > a {
		a = b
	}
	return a
}

// Returns the minimum of two ints
func Min(a, b int) int {
	if b < a {
		a = b
	}
	return a
}

func CalcArraySum(a []int) int {
	t := 0
	for _, v := range a {
		t += v
	}
	return t
}
