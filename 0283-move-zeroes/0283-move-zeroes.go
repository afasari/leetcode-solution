func moveZeroes(nums []int)  {
    l := len(nums)

    if l == 0{
        return
    }

    twoPointer(nums)
    
    // var count, idx int
    // for count != l{
    //     if nums[idx] == 0{
    //         nums = append(nums[:idx], nums[idx+1:]...)
    //         nums = append(nums, 0)
    //     } else {
    //         idx++
    //     }
    //     count++
    // }
}

func twoPointer(nums []int){
    left := 0
    for right := 0; right < len(nums); right++{
        if nums[right] != 0{
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }

    }
}