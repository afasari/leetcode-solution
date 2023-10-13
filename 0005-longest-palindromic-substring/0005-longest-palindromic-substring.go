func longestPalindrome(s string) string {
	var maxSubstring string
    n := len(s)
    
    for i := 0; i < n; i++ {
        for j := i; j <= n; j++ {
            substring := s[i:j]
			if isPalindrome(substring) {
                if len(maxSubstring) < len(substring){
					maxSubstring = substring
				}
			}
		}
	}
	return maxSubstring
}

func isPalindrome(s string) bool {
    n := len(s)
    mid := (n / 2)
    
	if n%2 == 0 {
		mid -= 1
	}
    
    for i := 0; i <= mid; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}