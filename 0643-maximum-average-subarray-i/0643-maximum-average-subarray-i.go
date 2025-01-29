import (
    "math"
)

func findMaxAverage(nums []int, k int) float64 {
    currentVal := 0
    maxVal := math.MinInt

    for i, num := range nums{
        currentVal += num

        if k-1 < i {
            currentVal -= nums[i-k]
        } 
        
        if k-1 <= i {
            if currentVal > maxVal{
                maxVal = currentVal
            }
        }

    }

    return float64(maxVal)/ float64(k)
}