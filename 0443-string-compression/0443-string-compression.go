func compress(chars []byte) int {
    lc := len(chars)
    if lc == 0{
        return lc
    }

    idx := 0
    i := 0
    for i < lc{
        ch := chars[i]
        count := 0
        for i < lc && chars[i] == ch {
            count++
            i++
        }

        chars[idx] = ch
        idx++

        if count != 1 {
            for _, digit := range[]byte(strconv.Itoa(count)) {
                chars[idx] = digit
                idx++
            }
        }
    }

    return idx
}