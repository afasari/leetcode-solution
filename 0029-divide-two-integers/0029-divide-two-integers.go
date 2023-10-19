func divide(dividend int, divisor int) int {
    
    dividend, isRet := checkVal(dividend)
    if isRet{
        return dividend
    }
    
    val := int(dividend/divisor)
    val, _ = checkVal(val)
    return val
}

func checkVal(val int) (int, bool){
    pow := int(math.Pow(2, 31))
    if val > 0 {
        if val > pow -1 {
            return pow-1, true
        }
    } else{
        if val < -pow{
            return -pow, true
        }
    }
    
    return val, false
}