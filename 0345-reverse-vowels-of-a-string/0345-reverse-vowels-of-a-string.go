func reverseVowels(s string) string {
    mapVowel := map[string]bool{
        "A": true,
        "I": true,
        "U": true,
        "E": true,
        "O": true,
    }

    // return rev(s, mapVowel)
    return reversed(s)

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

func isVowel(c rune) bool {
    // alternatively, we can just check 
    // return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || 
    //        c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
    c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func reversed(s string) string {
    ss := []rune(s)
    n := len(ss)
    l, r := 0, n - 1
    for l < r {
        // find the index of the first vowel in the given range
        for l < r && !isVowel(ss[l]) {
            l += 1
        }
        // find the index of last vowel in the given range
        for r > l && !isVowel(ss[r]) {
            r -= 1
        }
        // swap ss[l] and ss[r]
        ss[l], ss[r] = ss[r], ss[l]
        // since we've processed the character ss[l],
        // we move the left pointer to the right
        l += 1
        // since we've processed the character ss[r],
        // we move the right pointer to the left
        r -= 1
    }
    return string(ss)
}