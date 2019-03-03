func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	var path string
	dfsletterCombinations(digits, 0, path, &result)
	return result
}

func dfsletterCombinations(digits string, index int, path string, result *[]string) {
	Table := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if index == len(digits) {
		*result = append(*result, path)
		return
	} else {
		num := digits[index] - byte('0')
		for _, c := range Table[num] {
			path += string(c)
			dfsletterCombinations(digits, index+1, path, result)
			path = path[:len(path)-1]
		}
	}
}