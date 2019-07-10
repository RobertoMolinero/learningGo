package mystr

import "strings"

// Cat A self-written function for connecting multiple strings
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Join Call of the method implemented in the sdk to connect strings
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
