func maxArea(height []int) int {
    max := 0
    lh := len(height)
    if lh <2{
        return 0
    }
    
    left := 0
    right := lh-1
    for left < right{
        var area int
        if height[left] > height[right]{
            area = height[right] * (right-left)
            right--
        } else{
            area = height[left] * (right-left)
            left++
        }
        if area > max{
            max = area
        }
    }
    return max
}