func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	maxA, start, end := 0, 0, len(height)-1
	for start < end {
		maxA = max(maxA, min(height[start], height[end])*(end-start))
		if height[start] <= height[end] {
			start++
		} else {
			end--
		}
	}

	return maxA
}