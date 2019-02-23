func twoSum(nums []int, target int) []int {
	res := [2]int
	tbl := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := tbl[target-nums[i]]; ok {
			res[0] = val
			res[1] = i
			break
		} else {
			tbl[nums[i]] = i
		}
	}

	return res
}
