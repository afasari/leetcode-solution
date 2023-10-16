class Solution {
    /**
     * @param Integer[] $height
     * @return Integer
     */
    function maxArea($height) {
        
        $count = count($height);
        $r = 0;
        $i = 0;
        $j = $count - 1;
        
        while ($i < $count - 1)
        {    
            if ($height[$i] <= $height[$j])
            {
                $t = $height[$i] * ($j - $i);
                $i++;
            }
            else
            {
                $t = $height[$j] * ($j - $i);
                $j--;
            }

            if ($t > $r) 
            {
                $r = $t;
            }
        }
        
        return $r;
    }
}