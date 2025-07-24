func merge(intervals [][]int) [][]int {
    if len(intervals) < 2 {
        return intervals
    }

    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })


    res := [][]int{intervals[0]}
    for _, interval := range intervals[1:] {
        lastMerged := res[len(res)-1]
        if interval[0] <= lastMerged[1] {
            if interval[1] >= lastMerged[1] {
                lastMerged[1] = interval[1]
            }
        } else {
            res = append(res, interval)
        }
    }


    return res
}