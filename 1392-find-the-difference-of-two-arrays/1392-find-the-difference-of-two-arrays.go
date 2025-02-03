func findDifference(nums1 []int, nums2 []int) [][]int {
    m1 := make(map[int]bool)
    m2 := make(map[int]bool)
    md := make(map[int]bool)


    for idx := range nums1{
        m1[nums1[idx]] = true
    }

    for idx := range nums2{
        m2[nums2[idx]] = true
    }


    ret := make([][]int, 2)
    for n1, _ := range m1{
        if _, ok:= m2[n1]; ok{
            delete(m2, n1)
            delete(m1, n1)
            md[n1] = true
        }
    }

    for n2, _ := range m2{
        if _, ok:= m1[n2]; ok{
            delete(m2, n2)
            delete(m1, n2)
        }

        if _, ok:=md[n2]; ok{
            delete(m2, n2)
        }
    }

    fmt.Println(m1,m2,md)

        for n1, _ := range m1{
            ret[0] = append(ret[0], n1)
        }
        
        for n2, _ := range m2{
            ret[1] = append(ret[1], n2)
        }


    
    // for _, n2 := range nums2{
    //     if val, ok := md[n2]; ok && val == 1{
    //         ret[1] = append(ret[1], n2)
    //     }
    // }

    return ret
}