package aocutils

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Modulo(n int, m int) int {
	dividend := n

	if dividend < 0 {
		for dividend < 0 {
			dividend += m
		}
		return dividend
	}

	return dividend % m
}
