func strStr(haystack string, needle string) int {
    lenNeedle := len(needle)
    lenHaystack := len(haystack)
    if lenHaystack == lenNeedle{
        if haystack == needle{
            return 0
        }
        return -1
    }
    
    for i:=0; i<=lenHaystack-lenNeedle;i++{
        if haystack[i:i+lenNeedle] == needle{
            return i
        }
    }
    return -1
}