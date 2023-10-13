class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        // $len = strlen($s);
        // $map = array();
        // $max = 0;
        // for($i = 0; $i < $len; $i++){
        //     $val = $s[$i];
        //     if(isset($map[$val])){
        //         $exist_at = $map[$val];
        //         $map = array_slice($map, $exist_at, count($map)-1);
        //         $c = count($map);
        //         echo "idx $c $i $exist_at".PHP_EOL;
        //     } else {
                
        //     }
        //         $map[$val] = $i;
        //         echo "$i".PHP_EOL;
        //         if (count($map) > $max){
        //             $max = count($map);
        //         }
        //         print_r($map);
        // }
        // return $max;
        
        if (strlen($s) === 0) return 0;
        if (strlen($s) === 1) return 1;
        
        $chars = str_split($s);
        
        $i = $j = $max = 0;
        $seen = [];
        
        while ($i < count($chars))
        {
            $c = $chars[$i];
           
            while (array_key_exists($c, $seen))
            {
                unset($seen[$chars[$j]]);
        		$j++;
            }
            
            $seen[$chars[$i]] = true;
            
            $max = max($i - $j + 1, $max);
            $i++;
        }
        return $max;

    }
}