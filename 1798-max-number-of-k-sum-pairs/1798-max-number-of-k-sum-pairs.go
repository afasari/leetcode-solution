func maxOperations(nums []int, k int) int {
    count := 0

    m := make(map[int]int, 0)

    for _, num := range nums {
        if val, ok := m[num]; ok{
            if val == 1 {
                delete(m, num)
            } else {
                m[num]--
            }
            count++
        } else {
            m[k-num]++
        }
    }
    // l := 0
    // r := len(nums)-1

    // for l < r {
    //     nl := nums[l]
    //     nr := nums[r]

        
    // }


    return count
    
}