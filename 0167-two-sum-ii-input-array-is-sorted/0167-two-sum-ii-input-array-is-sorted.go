func twoSum(numbers []int, target int) []int {
    // O(n) two pointer
    l, r := 0, len(numbers)-1
    for r > l{
        count := numbers[l] + numbers[r]

        if count > target {
            r--
        } else if count < target {
            l++
        } else {
            return []int{l+1, r+1}
        }
    }

    return []int{}

    // O(n) hashmap
    // m := make(map[int]int)

    // for i, num := range numbers{
    //     if num > target {
    //         return []int{}
    //     }
    //     if j, ok := m[target-num]; ok{
    //         return []int{j+1,i+1}
    //     } else{
    //         m[num] = i
    //     }
    // }

    // O(N2)
    // for i:=0; i < len(numbers) -1; i++{
    //     for j:= 1; j < len(numbers);j++{
    //         if i!= j && numbers[i] + numbers[j] == target{
    //             return []int{i+1,j+1}
    //         }
    //     }
    // }

    return []int{}
}