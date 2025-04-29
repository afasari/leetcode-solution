func trap(height []int) int {
    var lmax, rmax, water, left int = 0, 0, 0, 0
    var right int = len(height) -1;

    for left < right {
        rmax = max(height[right], rmax);
        lmax = max(height[left], lmax);
        minH := min(lmax, rmax)

        if minH > height[left] {
            water += minH - height[left]
        }
        if minH > height[right] {
            water += minH - height[right]
        }

        if height[left] <= height[right] {
            left++;
        } else {
            right--;
        }
    }

    return water
}