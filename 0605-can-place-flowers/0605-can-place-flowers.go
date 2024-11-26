func canPlaceFlowers(flowerbed []int, n int) bool {
    // greedy
    // O(N)
    // O(1)
    m := len(flowerbed)
    total := 0
    for i := 0; i < m; i++ {
        if flowerbed[i] == 0 {
            if (i == 0 || flowerbed[i - 1] == 0) && (i == m - 1 || flowerbed[i + 1] == 0) {
                total++
                flowerbed[i] = 1
            }
        }
    }
    return total >= n
}