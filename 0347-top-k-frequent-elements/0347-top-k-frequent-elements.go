func topKFrequent(nums []int, k int) []int {
    var ret []int

    m := make(map[int]int)
    for _, num := range nums{
        m[num]++
    }

    mapMost := make(map[int][]int)
    var sliceCount []int

    for num, count := range m{
        mapMost[count] = append(mapMost[count], num)
        sliceCount = append(sliceCount, count)
        // sliceCount = append([]int{count}, sliceCount...) 
    }
    // sort.Ints(sliceCount)
    sort.Sort(sort.Reverse(sort.IntSlice(sliceCount)))

fmt.Println(sliceCount)
fmt.Println(mapMost)

    lastCount := 0
    for _, count := range sliceCount {
        fmt.Println(count)
        if count == lastCount{
            continue
        }
        lastCount = count
        if vals, ok := mapMost[count]; ok{
                    // fmt.Println(vals)
        sort.Ints(vals)

            ret = append(ret, vals...)
        }

        if len(ret) > k{
            return ret[:k]
        }
    }

    if len(ret) < k {
        return ret
    }
    
    return ret[:k]
}