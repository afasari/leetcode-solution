func productExceptSelf(nums []int) []int {
    l := len(nums)    
    if l == 0{
        return nums
    }

    ret := slices.Repeat([]int{1}, l)

    pref := 1
    post := 1
    
    for i := 0; i < l; i++{
        ret[i] = pref
        pref *= nums[i]
    }

    for i := l-1; i>=0; i-- {
        ret[i] *= post
        post *= nums[i]
    }

    return ret
}