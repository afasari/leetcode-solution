func getRow(rowIndex int) []int {
    ans := []int{1}
    prev := 1
    for i := 1; i <= rowIndex; i++{
        next := prev * (rowIndex - i + 1) / i
        ans = append(ans, next)
        prev = next
    }
    return ans
}