func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	start, cnt, max_cnt := 0, 0, 0
	tbl := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		tbl[[]rune(s)[i]]++
		cnt++

		//adjust start ptr until duplicate rune removed
		for tbl[[]rune(s)[i]] > 1 {
			tbl[[]rune(s)[start]]--
			start++
			cnt--
		}

		//update max counts
		if cnt > max_cnt {
			max_cnt = cnt
		}
	}
	return max_cnt
}