func plusOne(digits []int) []int {
    ld := len(digits)
    
    if digits[ld-1] <9 {
        digits[ld-1] +=1
    } else {
        digits[ld-1] = 0
        
        if ld == 1{
            return append([]int{1}, digits...)
        }
        
        for i:= ld-2; i >= 0; i--{
            inc := digits[i] 
            
            if inc == 9{
                digits[i] = 0
            } else{
                digits[i] += 1
                break
            }
            if i == 0{
                digits = append([]int{1}, digits...)
            }
        }
        
    }
    
    return digits
    
    
}