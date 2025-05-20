func findMaxAverage(nums []int, k int) float64 {
	var maxSum int
	var sum int
	for i := range k {
		sum += nums[i]
	}
	maxSum = sum
	for i := 1; i < len(nums)-k+1; i++ {
		sum = sum - nums[i-1] + nums[i+k-1]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}