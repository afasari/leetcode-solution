func isAnagram(s string, t string) bool {
    /*
    Sorting
    Time O(n log n + m log m)
    Space O(1) (other sort may different)
    */
    // if len(s) != len(t) {
    //     return false
    // }

    // sRunes, tRunes := []rune(s), []rune(t)
    // sort.Slice(sRunes, func(i, j int) bool {
    //     return sRunes[i] < sRunes[j]
    // })

    // sort.Slice(tRunes, func(i, j int) bool {
    //     return tRunes[i] < tRunes[j]
    // })

    // return string(sRunes) == string(tRunes)

    /*
    Hash Map
    Time O(n + m)
    Space O(1)
    */
    if len(s) != len(t) {
        return false
    }

    countS, countT := make(map[rune]int), make(map[rune]int)
    for i, ch := range s {
        countS[ch]++
        countT[rune(t[i])]++
    }

    for k, v := range countS {
        if countT[k] != v {
            return false
        }
    }

    return true

    /*
    Hash Table (array)
    Time O(n + m)
    Space O(1)
    */
    // if len(s) != len(t) {
    //     return false
    // }

    // count := [26]int{}
    // for i := 0; i < len(s); i++ {
    //     count[s[i]-'a']++
    //     count[t[i]-'a']--
    // }

    // for _, val := range count {
    //     if val != 0 {
    //         return false
    //     }
    // }

    // return true
}