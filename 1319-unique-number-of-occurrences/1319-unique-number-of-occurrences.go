func uniqueOccurrences(arr []int) bool {
    mOccur := make(map[int]int)

    for _, n := range arr{
        mOccur[n]++
    }

    if len(mOccur) == 1 {
        return true
    }

    mTotal := make(map[int]bool)
    for _,occur := range mOccur{
        if _, ok := mTotal[occur]; ok{
            return false
        } else {
            mTotal[occur] = true
        }

    }

    return true
    
}