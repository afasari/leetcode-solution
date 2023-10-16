class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[][]
     */
    function fourSum($nums, $target) {
                $ans = [];
        sort($nums);

        for($a=0; $a<count($nums)-3; $a++){
            if($a !== 0 && $nums[$a] === $nums[$a-1]){
                continue;
            }

            for($b=$a+1; $b<count($nums)-2; $b++){
                if($b !== $a+1 && $nums[$b] === $nums[$b-1]){
                    continue;
                }

                $d = $b + 1;
                $n = count($nums) - 1;
                while($d < $n){
                    if($d !== $b+1 && $nums[$d] === $nums[$d-1]){
                        $d++;
                        continue;
                    }

                    if($n !== count($nums) -1 && $nums[$n] === $nums[$n+1]){
                        $n--;
                        continue;
                    }
                    
                    $sum = $nums[$a] + $nums[$b] + $nums[$d] + $nums[$n];
                    
                    if($sum === $target){
                        $ans[] = [$nums[$a], $nums[$b], $nums[$d], $nums[$n]];
                        $d++;
                    }elseif($sum < $target){
                        $d++;
                    }else{
                        $n--;
                    }
                }
            }
        }

        return $ans;
    }
}