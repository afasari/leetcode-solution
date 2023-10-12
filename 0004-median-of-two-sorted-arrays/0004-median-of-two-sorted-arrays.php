class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Float
     */
    function findMedianSortedArrays($nums1, $nums2) {
        $nums3 = array_merge($nums1, $nums2);
        sort($nums3); 
        $count = count($nums3);
        if ($count %2 == 1){
            echo "asdf";
            print_r($nums3);
            return $nums3[$count/2];
        } else{
            echo "def";
            print_r($count);            
            return ($nums3[intdiv($count, 2) - 1] + $nums3[intdiv($count, 2)]) / 2; 
        }
    }
}