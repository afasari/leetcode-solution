func threeSum(nums []int) [][]int { // Returning [][]int
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// 1. Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 2. Standard Two-Pointer approach
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 3. Skip duplicates for left and right pointers
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}
