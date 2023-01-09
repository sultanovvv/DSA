package leetcode

func mySqrt(x int) int {
	return calcSqrt(1, x, x)
}

func calcSqrt(l int, h int, x int) int {
	mid := (l + h) / 2

	if mid == l {
		return l
	}

	if (mid)*(mid) > x {
		h = mid
	} else {
		l = mid
	}

	return calcSqrt(l, h, x)
}
