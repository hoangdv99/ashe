func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0
    if len(flowerbed) == 1 && flowerbed[0] == 0 {
        return 1 >= n && n >= 0
    }
    for i := 0; i < len(flowerbed); i++ {
        if (i == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0) || (i == len(flowerbed)-1 && flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0) {
            flowerbed[i] = 1
            count++
        }
        if flowerbed[i] == 0 && i-1>=0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
            flowerbed[i] = 1
            count++
        }
    }
    return n <= count
}