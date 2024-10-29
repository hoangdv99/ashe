
func singleNumber(nums []int) int {
	singles := []int{}
	for i := 0; i < len(nums); i++ {
		if exist(singles, nums[i]) {
			singles = remove(singles, nums[i])
		} else {
			singles = append(singles, nums[i])
		}
	}
	return singles[0]
}

func remove(nums []int, value int) []int {
	result := []int{}
	for _, v := range nums {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func exist(nums []int, value int) bool {
	for _, v := range nums {
		if v == value {
			return true
		}
	}
	return false
}