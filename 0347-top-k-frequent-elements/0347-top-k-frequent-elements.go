func topKFrequent(nums []int, k int) []int {
    // 1. Count frequencies: O(n)
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num]++
    }

    // 2. Create buckets where index = frequency: O(n)
    // bucket[i] contains all numbers that appear i times
    buckets := make([][]int, len(nums)+1)
    for num, freq := range counts {
        buckets[freq] = append(buckets[freq], num)
    }

    // 3. Iterate backwards from highest frequency: O(n)
    result := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)
        }
    }

    // Return exactly k elements (in case bucket contained more than needed)
    return result[:k]
}