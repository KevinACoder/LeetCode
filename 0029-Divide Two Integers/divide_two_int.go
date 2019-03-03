func divide(dividend int, divisor int) (result int) {
	var sign int
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	} else {
		sign = 1
	}

	if dividend < 0 {
		dividend *= -1
	}
	if divisor < 0 {
		divisor *= -1
	}

	for dividend >= divisor {
		dividend -= divisor
		result++
	}

	return sign * result
}