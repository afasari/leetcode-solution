import (
    // "unicode" 
    "strings"
)

func isChar(b byte) bool{
    return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
    // return unicode.IsLetter(rune(b)) || unicode.IsNumber(rune(b))
}
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right{
        for left<right && !isChar(s[left]){            
            left++
        }

        for left<right &&  !isChar(s[right]){
            right--
        }

        if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])){
            return false
        }

        left++
        right--
    }

    return true
}