// Brute-force
// func twoSum(nums []int, target int) []int {
//     var result []int

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				result = []int{i, j}
// 				break
// 			}
// 		}
// 	}

// 	return result
// }

// Hashmap
func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIndex, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIndex, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return nil
}