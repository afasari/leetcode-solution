func merge(nums1 []int, m int, nums2 []int, n int)  {
    for m > 0 || n > 0 {
        if m > 0 && (n == 0 || nums1[m-1] >= nums2[n-1]) {
            nums1[m+n-1] = nums1[m-1]
            m--
            continue
        }
        nums1[m+n-1] = nums2[n-1]
        n--
    }
}

// Time: O((n+m)log(n+m)
// Space: O(log(n+m))
// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     // Merge nums1 and nums2
//     nums1 = append(nums1[:n], nums2...)
    
//     // Sort nums1
//     sort.Ints(nums1)
// }

// Time: O((n+m)log(n+m)
// Space: O(log(n+m))
// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     sort.Ints(append(nums1[:n], nums2...))
// }

// TIme: O(ptr3) = O(m+n-1) = O(m+n)
// Spac: O(1)
// func merge(nums1 []int, m int, nums2 []int, n int)  { 
//     var ptr1, ptr2, ptr3 int = m-1, n-1, m+n-1
//     for ; ptr1 >= 0 && ptr2 >= 0; ptr3-- {
//         if nums2[ptr2] >= nums1[ptr1] {
//             nums1[ptr3] = nums2[ptr2]
//             ptr2--
//         } else {
//             nums1[ptr3] = nums1[ptr1]
//             ptr1--
//         }
//     }
    
//     if ptr2 >= 0 {
//         copy(nums1[:ptr3+1], nums2[:ptr2+1])
//     }
// }