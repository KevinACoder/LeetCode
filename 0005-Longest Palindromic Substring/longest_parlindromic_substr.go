func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	maxStart, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		start, end := i, i
		curLen := 1
		start--
		end++
		for start >= 0 && end < n && s[start] == s[end] {
			start--
			end++
			curLen += 2
		}

		if curLen > maxLen {
			maxStart = i - curLen/2
			maxLen = curLen
		}

		start = i
		end = i + 1
		curLen = 0
		for start >= 0 && end < n && s[start] == s[end] {
			start--
			end++
			curLen += 2
		}

		if curLen > maxLen {
			maxStart = i - curLen/2 + 1
			maxLen = curLen
		}
	}

	return s[maxStart:(maxStart + maxLen)]
}