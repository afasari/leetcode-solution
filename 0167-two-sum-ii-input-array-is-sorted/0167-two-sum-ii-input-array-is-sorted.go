func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)

    for i, num := range numbers{
        if j, ok := m[target-num]; ok{
            return []int{j+1,i+1}
        } else{
            m[num] = i
        }
    }
    // for i:=0; i < len(numbers) -1; i++{
    //     for j:= 1; j < len(numbers);j++{
    //         if i!= j && numbers[i] + numbers[j] == target{
    //             return []int{i+1,j+1}
    //         }
    //     }
    // }

    return []int{}
}