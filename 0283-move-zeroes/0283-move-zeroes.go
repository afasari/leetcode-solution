func moveZeroes(nums []int)  {
    l := len(nums)

    if l == 0{
        return
    }
    
    var count, idx int
    for count != l{
        if nums[idx] == 0{
            nums = append(nums[:idx], nums[idx+1:]...)
            nums = append(nums, 0)
        } else {
            idx++
        }
        count++
    }
}