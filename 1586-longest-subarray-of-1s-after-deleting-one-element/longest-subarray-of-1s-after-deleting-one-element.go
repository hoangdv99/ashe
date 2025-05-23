func longestSubarray(nums []int) int {
	start := 0
    max := 0
    k := 1
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			k--
		}

		for k < 0 {
			if nums[start] == 0 {
				k++
			}
			start++
		}

        if n := end - start + 1; n > max {
			max = n
		}
	}

	return max - 1
}