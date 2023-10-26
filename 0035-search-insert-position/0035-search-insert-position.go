func searchInsert(nums []int, target int) int {
//     ln := len(nums)
//     if target == nums[ln/2]{
//         return ln/2
//     }
    
//     if nums[ln/2] > target {
//         for i := ln/2 + 1; i < ln; i++{
//             if target == nums[i]{
//                 return i
//             }
//         }
//     } else{
//         for i := ln/2 + 1; i >= 0; i--{
//             if target == nums[i]{
//                 return i
//             }
//         }
//     }
    
//     return 0
    left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			right = mid
		default:
			left = mid + 1
		}
	}
	if right == len(nums)-1 && nums[right] < target {
		return right + 1
	}
	if right < 0 {
		return 0
	}
	return right
    
	// l := 0
	// r := len(nums) - 1
	// for l <= r {
	// 	m := l + (r-l)/2
	// 	v := nums[m]
	// 	switch {
	// 	case v < target:
	// 		l = m + 1
	// 	case v > target:
	// 		r = m - 1
	// 	default:
	// 		return m
	// 	}
	// }
	// return l
    
}