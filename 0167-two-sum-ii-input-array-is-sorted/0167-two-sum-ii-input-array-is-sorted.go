func twoSum(numbers []int, target int) []int {
    for i:=0; i < len(numbers) -1; i++{
        for j:= 1; j < len(numbers);j++{
            if i!= j && numbers[i] + numbers[j] == target{
                return []int{i+1,j+1}
            }
        }
    }
    
    return []int{}
}