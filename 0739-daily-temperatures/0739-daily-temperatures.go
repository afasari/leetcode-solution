func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    // Monotonic stack to store indices
    stack := []int{}

    for i, currTemp := range temperatures {
        // While stack is not empty AND current temp is warmer than stack top index's temp
        for len(stack) > 0 && currTemp > temperatures[stack[len(stack)-1]] {
            // Pop the index
            prevIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            // Calculate the distance (wait time)
            ans[prevIndex] = i - prevIndex
        }
        // Push current index onto stack
        stack = append(stack, i)
    }

    return ans
}