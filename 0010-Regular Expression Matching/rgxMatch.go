func isMatch(str string, pattern string) bool {
	n, m := len(str), len(pattern)
	if n == 0 {
		return m == 0
	} else if m == 0 {
		return false
	}

	lookup := [n + 1][m + 1]bool{}
	lookup[0][0] = false
	for j := 1; j <= m; j++ {
		if pattern[j-1] == '*' {
			lookup[0][j] = lookup[0][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if pattern[j-1] == '*' {
				lookup[i][j] = lookup[i-1][j] || lookup[i][j-1]
			} else if pattern[j-1] == '.' || pattern[j-1] == str[i-1] {
				lookup[i][j] = lookup[i-1][j-1]
			} else {
				lookup[i][j] = false
			}
		}
	}

	return lookup[n][m]
}