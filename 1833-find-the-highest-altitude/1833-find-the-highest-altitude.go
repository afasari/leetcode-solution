func largestAltitude(gain []int) int {
    var curHeight, maxHeight int

    for _, g := range gain{
        curHeight += g
        if curHeight > maxHeight{
            maxHeight = curHeight
        }
    }

    return maxHeight
    
}