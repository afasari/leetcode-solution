class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        return $this->oneloophash($nums, $target);
    }

    function bruteforce($nums, $target){
        $len = count($nums);
        for ($i = 0; $i < $len-1; $i++){
            for ($j=$i+1; $j<$len;$j++){
                if ($nums[$i] + $nums[$j] === $target){
                    return [$i,$j];
                }
            }
        }
    }

    function oneloophash($nums, $target){
        $map = array();
        $len = count($nums);

        for ($i = 0; $i < $len; $i++){
            $difference = $target-$nums[$i];
            if (array_key_exists($difference, $map)){
                return [$i, $map[$difference]];
            }
            $map[$nums[$i]] = $i;
        }
    }
    

    function twoloophash($nums, $target){
        $map = array();
        $len = count($nums);

        for ($i = 0; $i < $len; $i++){
            $map[$nums[$i]] = $i;
        }

        for ($i = 0; $i < $len; $i++){
            $difference = $target-$nums[$i];
            if (array_key_exists($difference, $map) && $map[$difference] != $i){
                return [$i, $map[$difference]];
            }
        }
    }
}