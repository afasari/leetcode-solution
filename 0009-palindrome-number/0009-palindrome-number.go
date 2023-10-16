func isPalindrome(x int) bool {
    def := x
    res := 0
    for x > 0{
        digit := x % 10
        res = res * 10
        res += digit
        x /= 10
    }

fmt.Println(res, x)
    return res == def
    
}