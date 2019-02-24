func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	str := make([]string, numRows)
	t := numRows*2 - 2 //every 2, 4, 6 item is gapped
	for i := 0; i < len(s); i++ {
		index := i % t
		if index < numRows {
			str[index] += string(s[i])
		} else {
			str[t-index] += string(s[i])
		}
	}

	var result string
	for _, st := range str {
		result += st
	}

	return result
}