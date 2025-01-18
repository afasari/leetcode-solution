func reverseVowels(s string) string {
    mapVowel := map[string]bool{
        "A": true,
        "I": true,
        "U": true,
        "E": true,
        "O": true,
    }

    ret := ""
    reversedVowel := ""
    mapRet := make(map[int]string)

    for idx, b := range s{
        strB := string(b)
        if _, ok := mapVowel[strings.ToUpper(strB)]; ok{
            reversedVowel = strB + reversedVowel
        } else {
            mapRet[idx] = strB
        }
    }

    idxVowel := 0
    for idx, _ := range s{
        strRet := ""
        if val, ok := mapRet[idx]; ok {
            strRet = val
        } else {
            strRet = string(reversedVowel[idxVowel])
            idxVowel++
        }
        ret += strRet
    }
    
    return ret
}