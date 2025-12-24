func reverseBits(n int) int {
    bin := fmt.Sprintf("%032b",n)

    ret := ""
    for i := len(bin)-1; i >= 0 ; i-- {
        ret += fmt.Sprintf("%v",string(bin[i]))
    }

    retInt, _ := strconv.ParseInt(ret, 2, 64)
    return int(retInt)
}