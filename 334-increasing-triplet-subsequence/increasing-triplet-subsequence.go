func increasingTriplet(nums []int) bool {
    first := math.MaxInt
    second := math.MaxInt
    for _, v := range nums {
        if v <= first {
            first = v
        } else if v <= second {
            second = v
        } else {
            return true
        }
    }
    return false
}