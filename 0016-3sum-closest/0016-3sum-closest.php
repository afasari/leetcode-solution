class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function threeSumClosest($nums, $target) {
        $diff  = PHP_INT_MAX;
        $n = count($nums);
        sort($nums);

        for($i = 0;$i <= $n;$i++){
            $low = $i + 1;
            $high = $n - 1;
            while($low < $high) {
                $sum = $nums[$i] + $nums[$low] +$nums[$high];
                if(abs($target- $sum) < abs($diff)) {
                    $diff = $target- $sum;
                }
                if($sum < $target) {
                    $low++;
                } else {
                    $high--;
                }
            }
        }
        return $target - $diff;
    }
}