func shuffle(nums []int, n int) []int {
    var ret []int
    ln := len(nums)
    for i:=0; i < ln/2; i ++ {
        ret = append(ret, nums[i])
        ret = append(ret, nums[i+(ln/2)])
    }

    return ret
}