package gocod

import "errors"

// ertweiterter euklidischer algorithmus
func EEA(x, y int) (ggt, s, t int, err error) {
	if x <= 0 || y <= 0 {
		return 0, 0, 0, errors.New("x and y must be greater than 0")
	}
	if x < y {
		x, y = y, x
	}
	var r, s1, s2, t1, t2 int = x, 1, 0, 0, 1
	for r != 0 {
		q := x / y
		r = x % y
		x, y = y, r
		s = s1 - q*s2
		s1, s2 = s2, s
		t = t1 - q*t2
		t1, t2 = t2, t
	}
	ggt = x
	s = s1
	t = t1
	return
}
