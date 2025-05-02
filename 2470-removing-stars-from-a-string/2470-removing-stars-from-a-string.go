func removeStars(s string) string {
    ret := make([]byte, len(s))
    idx := 0
    
    for i, char := range s {
        if char != '*' {
            ret[idx] = s[i]
            idx++
        } else {
            ret[idx] = 0
            idx--
        }
    }

    return string(ret[:idx])
}