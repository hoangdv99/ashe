// Approach 1: Use 2 loop to calculate each area and find the max area
// func maxArea(height []int) int {
// 	n := len(height)
// 	var max int
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			var h int
// 			if height[i] <= height[j] {
// 				h = height[i]
// 			} else {
// 				h = height[j]
// 			}
// 			area := h * (j - i)
// 			if area > max {
// 				max = area
// 			}
// 		}
// 	}
// 	return max
// }

// Approach 2: Use 2 pointers, one points to head of array and the second points to end of array. Compare 2 heights of 2 pointers, move to the lower. On each move, calculate the area and find the max

func maxArea(height []int) int {
	n := len(height)
	i := 0
	j := n - 1
	var max int
	for i < j {
		var h int
		h = min(height[i], height[j])
		area := h * (j - i)
		if area > max {
			max = area
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}