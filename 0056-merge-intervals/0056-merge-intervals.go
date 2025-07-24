func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    if len(intervals) < 2 {
        return intervals
    }

    res := [][]int{intervals[0]}
    for _, interval := range intervals[1:] {
        if interval[0] <= res[len(res)-1][1] {
            if interval[1] >= res[len(res)-1][1] {
                res[len(res)-1][1] = interval[1]
            }
        } else {
            res = append(res, interval)
        }
    }


    return res
}