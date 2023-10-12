func twoSum(nums []int, target int) []int {
    // return onePassHash(nums, target)
    // return bruteForce(nums, target)
    return twoPassHash(nums, target)
}

func bruteForce(nums []int, target int) []int{
    // time O(n2)
    // space O(1)
    for i:= 0; i < len(nums) - 1; i++{
        for j :=1 ; j < len(nums);j++{
            if i != j && nums[i] + nums[j] == target{
                return []int{i, j}
            }
        }
    }
    return []int{}
}

func twoPassHash(nums []int, target int) []int{
    // time O(n)
    // space O(n)
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
        indexMap[currNum] = currIndex
    }
    
    for i := 0; i < len(nums); i++{
        complement := target - nums[i]
        if idx, ok := indexMap[complement]; ok && idx != i{
            return []int{i, idx}
        }
    }
    return []int{}
}

func onePassHash(nums []int, target int) []int{
    // time O(n)
    // space O(n)
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}