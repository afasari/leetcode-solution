func pivotIndex(nums []int) int {

    pivot := -1

    sum := 0
    for _, num := range nums{
        sum += num
    }
    
    
    prefix := 0
    for idx, num := range nums{
        if (sum-num)%2 == 0 && prefix == (sum-num)/2{
            return idx
        } else {
            prefix += num
        }
    }

    return pivot
}