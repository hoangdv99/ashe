func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	left := 0
	for i, v := range nums {
		left += v
		if left-v == total-left {
			return i
		}
	}
	return -1
}