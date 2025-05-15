func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, v := range candies {
        if v > max {
            max = v
        }
    }
    result := make([]bool, len(candies))
    for i, v := range candies {
        if v + extraCandies >= max {
            result[i] = true
        } else {
            result[i] = false
        }
    }
    return result
}