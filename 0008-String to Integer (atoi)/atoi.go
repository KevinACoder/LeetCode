func myAtoi(str string) int {
	const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	const (
		MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
		MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
		MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
	)
	var res, sign int64
	sign = 1
	for _, s := range str {
		if s == '-' {
			sign = -1
		} else if s >= '0' && s <= '9' {
			res = res*10 + int64(s-'0')
		} else if s == ' ' || s == '+' {

		} else {
			return 0
		}

		if res > MaxInt {
			return 0
		}
	}

	return int(sign * res)
}