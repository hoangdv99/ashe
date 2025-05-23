func findDifference(nums1 []int, nums2 []int) [][]int {
	n := max(len(nums1), len(nums2))
	freq1 := make(map[int]bool)
	freq2 := make(map[int]bool)
	for i := 0; i < n; i++ {
		if i < len(nums1) {
			freq1[nums1[i]] = true
		}
		if i < len(nums2) {
			freq2[nums2[i]] = true
		}
	}

	res1 := []int{}
	res2 := []int{}
	for i := 0; i < n; i++ {
		if i < len(nums1) && !freq2[nums1[i]] {
			res1 = append(res1, nums1[i])
			freq2[nums1[i]] = true
		}
		if i < len(nums2) && !freq1[nums2[i]] {
			res2 = append(res2, nums2[i])
			freq1[nums2[i]] = true
		}
	}

	return [][]int{res1, res2}
}
