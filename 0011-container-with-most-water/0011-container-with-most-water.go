func maxArea(height []int) int {
    maxArea := 0
    left := 0
    right := len(height) - 1

    for left < right {
        // 1. Calculate current area
        w := right - left
        h := 0
        if height[left] < height[right] {
            h = height[left]
        } else {
            h = height[right]
        }
        
        currentArea := w * h
        
        // 2. Update global max
        maxArea = max(maxArea, currentArea)

        // 3. Move the pointer pointing to the shorter line
        // This is the greedy step to find a taller bottleneck
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}