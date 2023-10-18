class Solution {

    /**
     * @param String $haystack
     * @param String $needle
     * @return Integer
     */
    function strStr($haystack, $needle) {
        $lh = strlen($haystack);
        $ln = strlen($needle);
        
        if ($lh == $ln){
            if ($haystack == $needle){
                return 0;
            } 
            return -1;
        }
        for($i=0;$i<=$lh-$ln;$i++){
            if (substr($haystack,$i,$ln) == $needle){
                return $i;
            }
        }
        return -1;
    }
}