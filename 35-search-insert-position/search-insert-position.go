func searchInsert(nums []int, target int) int {
	for i, e := range nums {
		if e >= target {
			return i
		}
	}
	return len(nums)
}