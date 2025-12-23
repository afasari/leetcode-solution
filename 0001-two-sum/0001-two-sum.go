func twoSum(nums []int, target int) []int {
    /*
    Brute Force
    Time O(n^2)
    Space O(1)
    */
    // for i := 0; i < len(nums); i++ {
    //     for j := i + 1; j < len(nums); j++ {
    //         if nums[i] + nums[j] == target {
    //             return []int{i, j}
    //         }
    //     }
    // }
    // return []int{}
    
    /*
    Sorting
    Time O(n log n)
    Space O(n)
    */
    // A := make([][2]int, len(nums))
    // for i, num := range nums {
    //     A[i] = [2]int{num, i}
    // }

    // sort.Slice(A, func(i, j int) bool {
    //     return A[i][0] < A[j][0]
    // })

    // i, j := 0, len(nums)-1
    // for i < j {
    //     cur := A[i][0] + A[j][0]
    //     if cur == target {
    //         if A[i][1] < A[j][1] {
    //             return []int{A[i][1], A[j][1]}
    //         } else {
    //             return []int{A[j][1], A[i][1]}
    //         }
    //     } else if cur < target {
    //         i++
    //     } else {
    //         j--
    //     }
    // }
    // return []int{}

    /*
    Hash Map
    Time O(n)
    Space O(n)
    */
    //     indices := make(map[int]int)

    // for i, n := range nums {
    //     indices[n] = i
    // }

    // for i, n := range nums {
    //     diff := target - n
    //     if j, found := indices[diff]; found && j != i {
    //         return []int{i, j}
    //     }
    // }
    // return []int{}

    prevMap := make(map[int]int)

    for i, n := range nums {
        diff := target - n
        if j, found := prevMap[diff]; found {
            return []int{j, i}
        }
        prevMap[n] = i
    }
    return []int{}
}