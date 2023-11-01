func plusOne(digits []int) []int {
    ld := len(digits)
    
    if digits[ld-1] <9 {
        digits[ld-1] +=1
        return digits
    } 
        for i:= ld-1; i >= 0; i--{
            inc := digits[i] 
            
            if inc != 9{
                digits[i] += 1
                break
            }
            
            digits[i] = 0
            if i == 0{
                digits = append([]int{1}, digits...)
            }
        }
    
    return digits
    
    
}