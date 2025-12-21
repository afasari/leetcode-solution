func canMakeArithmeticProgression(arr []int) bool {
    la := len(arr)
    if la <= 2 {
        return true
    }

    sort.Ints(arr)
    diff := arr[1] - arr[0]
    for i := 1; i < la; i++ {
        if arr[i]- arr[i-1] != diff {
            return false
        }
    } 

    return true
}