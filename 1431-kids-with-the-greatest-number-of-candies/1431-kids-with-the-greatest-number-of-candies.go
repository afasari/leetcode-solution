func kidsWithCandies(candies []int, extraCandies int) []bool {
    
    ret := make([]bool, len(candies))
    
    // cara 1, O(2n)
    maxCandy := 0
    for _, candy := range candies{
        if candy > maxCandy {
            maxCandy = candy
        }
    }
    
    for i, candy := range candies{
        if candy + extraCandies >= maxCandy {
            ret[i] = true
        } 
    }
    
    return ret
}