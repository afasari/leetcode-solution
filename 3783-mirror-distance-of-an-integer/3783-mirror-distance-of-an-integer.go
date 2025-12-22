func mirrorDistance(n int) int {
    reverse := 0
    temp := n
    for temp != 0 {
        reverse = (reverse *10) + (temp % 10)
        temp /= 10
    }

    return int(math.Abs(float64(n-reverse)))
}