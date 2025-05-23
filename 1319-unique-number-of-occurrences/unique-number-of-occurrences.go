func uniqueOccurrences(arr []int) bool {
    freq := make(map[int]int)
    freq1 := make(map[int]int)
    for _, v := range arr {
        freq[v]++
    }
    for _, f := range freq {
        freq1[f]++
        if freq1[f] > 1 {
            return false
        }
    }
    return true
}