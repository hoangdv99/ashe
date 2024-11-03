func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
    for _, num := range nums {
        _, exist := m[num]
        if !exist {
            m[num] = 1
        } else {
            return true
        }
    }
    return false
}