func lastStoneWeight(stones []int) int {
    if len(stones) == 0 {
        return 0
    }
    
    if len(stones) == 1 {
        return stones[0]
    }

    if len(stones) == 2{
    return int(math.Abs(float64(stones[0]-stones[1])))
    }

   for {
        if len(stones) == 0 { return 0 }
        if len(stones) == 1 { return stones[0] }
        x, y := 0, 0
        index := 0
        for i, v := range stones {
            if v > y {
                y = v
                index = i
            }
        }
        stones = append(stones[:index], stones[index + 1:]...)
        for i, v := range stones {
            if v > x {
                x = v
                index = i
            }
        }
        stones = append(stones[:index], stones[index + 1:]...)
        if y != x { stones = append(stones, y - x) }
    }

    return -1
}