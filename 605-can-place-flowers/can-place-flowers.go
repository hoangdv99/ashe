func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0
    length := len(flowerbed)

    for i := 0; i < length; i++ {
        if flowerbed[i] == 0 {
            emptyLeft := (i == 0) || (flowerbed[i-1] == 0)
            emptyRight := (i == length-1) || (flowerbed[i+1] == 0)

            if emptyLeft && emptyRight {
                flowerbed[i] = 1
                count++
                if count >= n {
                    return true
                }
            }
        }
    }

    return count >= n
}