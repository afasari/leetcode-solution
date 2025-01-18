func reverseVowels(s string) string {
    mapVowel := map[string]bool{
        "A": true,
        "I": true,
        "U": true,
        "E": true,
        "O": true,
    }

    return rev(s, mapVowel)

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

func rev(s string, mapVowel map[string]bool) string{
    if len(s) == 1{
        return s
    }
    start := 0
    end := len(s)-1
    leftS := ""
    rightS := ""

    for start < end {
        _, isLeftVowel := mapVowel[strings.ToUpper(string(s[start]))]
        _, isRightVowel := mapVowel[strings.ToUpper(string(s[end]))]

        isEndLoop := end - start == 1

        if !isLeftVowel || (isEndLoop && !isRightVowel) {
            leftS += string(s[start])
            start++
        }
        
        if !isRightVowel || (isEndLoop && !isLeftVowel) {
            rightS = string(s[end]) + rightS
            end--
        }

        if isLeftVowel && isRightVowel{
            leftS += string(s[end])
            rightS = string(s[start]) + rightS
            start++
            end--
        }
        
        if start == end {
            leftS += string(s[start])
            break
        }
    }

    return leftS + rightS
}