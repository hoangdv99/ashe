// func majorityElement(nums []int) int {
//     m := make(map[int]int)
//     for _, num := range nums {
//         _, ok := m[num]
//         if !ok {
//             m[num] = 1
//         } else {
//             m[num] += 1
//         }
//         if m[num] > len(nums) / 2 {
//             return num
//         } 
//     }
//     return 0
// }

func majorityElement(nums []int) int {
    count := 0
    candidate := 0
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if candidate == num {
            count++
        } else {
            count--
        }
    }
    return candidate
}