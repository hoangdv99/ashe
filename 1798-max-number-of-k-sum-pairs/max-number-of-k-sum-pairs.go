func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	var count int
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] < k {
			i++
		} else if nums[i]+nums[j] == k {
			i++
			j--
			count++
		} else {
			j--
		}
	}
	return count
}