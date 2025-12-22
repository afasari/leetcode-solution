func isPalindrome(x int) bool {
    if x == 1 {
        return true
    }

    if x < 0 {
        return false
    }
    reversed := 0
    temp := x

    for temp != 0 {
        reversed = (reversed * 10) + (temp % 10)
        temp /= 10
    }

    // xs := strconv.Itoa(x)

    return reversed == x
}