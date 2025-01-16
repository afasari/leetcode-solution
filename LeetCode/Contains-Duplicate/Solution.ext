func containsDuplicate(nums []int) bool {
    hashMap := make(map[int]bool)

    for _, num := range nums{
        if _, ok := hashMap[num]; ok{
            return true
        }
        hashMap[num] = true
    }
    return false
}