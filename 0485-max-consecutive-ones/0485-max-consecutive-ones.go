func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    cm := 0

    for _, num := range nums{
        if num == 1{
            cm++
            if cm > max {
                max = cm
            }
        } else {
            cm = 0
        }
    }

    return max
}