
import {
	"math"
}
func reverse(x int) int {
	const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	const (
		MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
		MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
		MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
	)

	if x == int(math.MinInt32) || x == 0 {
		return 0
	}
	sign := int64(1)
	xl := int64(x)
	rl := int64(0)
	var digits []int64
	if xl < 0 {
		sign = -1
		xl *= xl
	}

	for xl > 0 {
		digits = append(digits, xl%10)
		xl /= 10
	}

	for _, d := range digits {
		rl = rl*10 + d //igits[i]
	}
	rl *= sign

	if rl > int64(MaxInt) || rl < int64(MinInt) {
		return 0
	} else {
		return int(rl)
	}
}