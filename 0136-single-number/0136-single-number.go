func singleNumber(nums []int) int {
    m := make(map[int]bool)
    for _, num := range nums{
        if val, ok := m[num]; ok && val{
            fmt.Println("masuk del", num)
            delete(m, num)
        } else {
         m[num] = true           
        }
    }
    
    if len(m) == 0{
        return -1
    }
    
    
    fmt.Println(m)
    for mv, _ := range m {
        return mv
    }
    
    return  -1 
}