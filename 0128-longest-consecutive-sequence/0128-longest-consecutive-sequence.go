func longestConsecutive(nums []int) int {
    var ret int

    mapNums := make(map[int]bool)
    
    for _, n := range nums{
        mapNums[n] = true
    }

    for n, _ := range mapNums{
        if mapNums[n-1]{
            continue
        }

        curLen := 0
        for mapNums[n]{
            curLen++
            n++
        }
        
        ret = max(ret, curLen)
    }

    return ret
}