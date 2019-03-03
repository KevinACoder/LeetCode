func generateParenthesis(n int) (result []string) {
	path := ""
	dfsgenerateParenthesis(&path, n, n, &result)
	return result
}

func dfsgenerateParenthesis(path *string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, *path)
		return
	}

	if left > 0 {
		*path += "("
		dfsgenerateParenthesis(path, left-1, right, result)
		*path = (*path)[:len(*path)-1]
	}

	if right > left { //make balance
		*path += ")"
		dfsgenerateParenthesis(path, left, right-1, result)
		*path = (*path)[:len(*path)-1]
	}
}