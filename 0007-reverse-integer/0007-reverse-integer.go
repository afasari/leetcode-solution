func reverse(x int) int {
    const MaxUint32 = ^uint32(0) 
    const MaxInt32 = int(MaxUint32 >> 1) 
    const MinInt32 = int(-MaxInt32- 1) 

    new_int := 0
    for x != 0 {
        remainder := x % 10
        new_int *= 10
        new_int += remainder 
        x /= 10
    }
    if new_int > MaxInt32 || new_int < MinInt32{
        return 0
    }

    return new_int 
}