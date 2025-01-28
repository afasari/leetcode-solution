func isSubsequence(s string, t string) bool {
    idx := 0
    if len(s) == 0{
        return true
    }

    for _, ts := range t{
        if idx == len(s){
            break
        }

        if s[idx] == byte(ts) {
            idx++
        }
    }


    return idx == len(s)
    
}