func closeStrings(word1 string, word2 string) bool {
    l1 := len(word1)
    l2 := len(word2)

    if l1 != l2 {
        return false
    }
    
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    for i := 0; i < l1; i++{
        m1[string(word1[i])]++
        m2[string(word2[i])]++
    }

    lm1 := len(m1)
    lm2 := len(m2)
    if lm1 != lm2 {
        return false
    }

    md1 := make(map[int]int)
    md2 := make(map[int]int)
    for word, count := range m1 {
        md1[count]++
        md2[m2[word]]++
    }

    for count, times := range md1 {
        if md2[count] != times{
            return false
        }
    }
    return true
}