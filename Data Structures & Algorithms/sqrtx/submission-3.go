func mySqrt(x int) int {
	
	if x < 2 {
		return x
	}

	left := 1
	right := x
	closest := 0

	for left <= right {
		mid := (left + right) / 2
		sqr := mid * mid

		if sqr == x {
			return mid
		} else if sqr < x {
			closest = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return closest
}