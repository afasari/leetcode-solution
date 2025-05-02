func maxVowels(s string, k int) int {
    ls := len(s)
    m := 0
    count := 0

    for i := 0; i < k; i++ {
        if isVowel(s[i]) {
            count++
        }
    }

    m = count

    for i := k; i < ls; i++ {
        if isVowel(s[i]) {
            count++
        }

        if isVowel(s[i-k]) {
            count--
        }

        m = max(m, count)

        if m == k {
            return m
        }
    }

    return m
}

func isVowel(s byte) bool{
    m := map[byte]bool{
        'a': true,
        'i': true,
        'u': true,
        'e': true,
        'o': true,
    }

    return m[s]
}