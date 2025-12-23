func containsDuplicate(nums []int) bool {
    /*
    Brute Force 
    Time O(n^2) 
    Space O(1)
    */
    // for i := 0; i < len(nums); i++ {
    //     for j := i + 1; j < len(nums); j++ {
    //         if nums[i] == nums[j] {
    //             return true
    //         }
    //     }
    // }
    // return false

    /*
    Sorting
    Time O(n log n)
    Space O(log n) (other sort may different)
    */
    // sort.Ints(nums)
    // for i := 1; i < len(nums); i++ {
    //     if nums[i] == nums[i-1] {
    //         return true
    //     }
    // }
    // return false

    /*
    Hash
    Time O(n)
    Space O(n)
    */
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}
