func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	totalWater := 0

	for left < right {
		// We can only trap water based on the shorter side
		if leftMax < rightMax {
			left++
			// Update leftMax
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				// Trapped water = boundary - current height
				totalWater += leftMax - height[left]
			}
		} else {
			right--
			// Update rightMax
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
		}
	}

	return totalWater
}
