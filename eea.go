package gocod

import "errors"

// erweiterter euklidischer algorithmus
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

// erweiterter euklidischer algorithmus
/*
function extended_gcd(a, b)
    (old_r, r) := (a, b)
    (old_s, s) := (1, 0)
    (old_t, t) := (0, 1)

    while r ≠ 0 do
        quotient := old_r div r
        (old_r, r) := (r, old_r − quotient × r)
        (old_s, s) := (s, old_s − quotient × s)
        (old_t, t) := (t, old_t − quotient × t)

    output "Bézout coefficients:", (old_s, old_t)
    output "greatest common divisor:", old_r
    output "quotients by the gcd:", (t, s)
*/
func EeaOpt(a, b int) (int, int, int, error) {
	var (
		old_r, r int = a, b
		old_s, s int = 1, 0
		old_t, t int = 0, 1
	)

	for r != 0 {
		quotient := old_r / r
		old_r, r = r, old_r-quotient*r
		old_s, s = s, old_s-quotient*s
		old_t, t = t, old_t-quotient*t
	}

	return old_r, old_s, old_t, nil
}
