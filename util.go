package gocolor

// Credit https://stackoverflow.com/questions/28058278/how-do-i-reverse-a-slice-in-go
func reverse(s []colorAttribute) []colorAttribute {
	a := make([]colorAttribute, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
