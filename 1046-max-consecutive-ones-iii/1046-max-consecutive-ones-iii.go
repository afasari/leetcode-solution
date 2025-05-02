func longestOnes(nums []int, k int) int {
    ln := len(nums)
    i, j := 0, 0
    for j < ln {
        if nums[j] == 0 {
            k--
        }
        j++
        if k < 0 {
            if nums[i] == 0 {
                k++
            }
            i++
        }
    }
    
    return j - i
}