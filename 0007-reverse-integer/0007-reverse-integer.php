class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function reverse($x) {
        // if ($x <10){
        //     return $x;
        // }
        // $mod = $x % 10;
        // return $mod*10+ $this->reverse(floor($x/10));
        
        $rev = 0;
        while ($x != 0) {
            $digit = $x % 10;
            $rev = $rev * 10 + $digit;
            $x = ($x < 0) ? ceil($x / 10) : floor($x / 10);
        }
        
        if ($rev < pow(-2, 31) || $rev > pow(2, 31) - 1) {
            return 0;
        }
        
        return $rev;
    }
}