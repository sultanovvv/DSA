package leetcode

func mySqrt(x int) int {
	n := x
	for {
		if (n)*(n) <= x {
			break
		}
		n = n - n/2
	}
	for {
		n++
		if (n)*(n) > x {
			n--
			break
		}
		if (n)*(n) == x {
			break
		}
	}
	return n
}
