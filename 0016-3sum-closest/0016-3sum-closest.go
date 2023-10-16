func threeSumClosest(nums []int, target int) int {
    if len(nums) < 3 {return -1}
    
    sort.Ints(nums)
    smallestDiff := 1 << 31
    
    for i := 0; i < len(nums) - 2; i++ {
        left, right := i + 1, len(nums) - 1
        
        for left < right {
            diff := target - nums[i] - nums[left] - nums[right]
            if diff == 0 {return target}
            
            if abs(diff) < abs(smallestDiff) || (abs(diff) == abs(smallestDiff) && diff > smallestDiff) {
                smallestDiff = diff
            }
            
            if diff > 0 {
                left++
            } else {
                right--
            }
        }
    }
    
    return target - smallestDiff
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    
    return -x
}