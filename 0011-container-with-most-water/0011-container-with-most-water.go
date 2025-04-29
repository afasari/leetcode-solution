func maxArea(height []int) int {
    l, water := 0, 0
    r := len(height)-1

    for l < r {
        hl := height[l]
        hr := height[r]
        minH := hl

        if hl > hr{
            minH = hr
        }

        temp := (r-l) * minH
        if temp > water {
            water = temp
        }

        if hl > hr {
            r--
        } else {
            l++
        }
    }

    return water
    
}