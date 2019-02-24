func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}

	return x == y || y\10 == x
}