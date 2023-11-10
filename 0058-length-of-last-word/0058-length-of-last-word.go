func lengthOfLastWord(s string) int {
    words := strings.Fields(s)

    lwords := len(words)
    if lwords == 0{
        return 0
    }
    return len(words[lwords-1])
}