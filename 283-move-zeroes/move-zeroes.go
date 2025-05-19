func moveZeroes(nums []int) {
	var tmp int
	for _, v := range nums {
		if v != 0 {
			nums[tmp] = v
			tmp++
		}
	}
	for i := tmp; i < len(nums); i++ {
		nums[i] = 0
	}

}