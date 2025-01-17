func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    mapWord1 := make(map[byte]int)
    mapWord2 := make(map[byte]int)
    length := len(s)
    for i:=0; i< length; i++{
        mapWord1[s[i]]++
        mapWord2[t[i]]++
    }

    for key, val := range mapWord1 {
        if val != mapWord2[key]{
return false
        }
    }

    return true
    
}