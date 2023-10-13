class Solution {

    /**
     * @param String $s
     * @return String
     */
    function longestPalindrome($s) {
        $s_rev = strrev($s);
        $i = 1;
        $str = $s[0];
        $res = $str;
        
        while($i < strlen($s)) {
            if (strpos($s_rev, $str) === false) $str = substr($str, 1);
            else {
                $str .= $s[$i];
                $i++;
            }
            if (strpos($s_rev, $str) !== false && strlen($res) < strlen($str) && $str == strrev($str)) $res = $str;
        }
        return $res;
        // if (($length = strlen($s)) <= 1) {
        //     return $s;
        // }
        // if (strrev($s) === $s) {
        //     return $s;
        // }

        // $max_length = 1;
        // for ($i = 0; $i < $length; ++$i) {
        //     for ($len = $max_length; $len <= $length; ++$len) {
        //         $start = $i - ($len >> 1);
        //         if ($start < 0 || $start + $len > $length) {
        //             break;
        //         }
        //         $substr = substr($s, $start, $len);
        //         if ($substr === strrev($substr)) {
        //             $str = $substr;
        //             $max_length = $len;
        //         } else if ($max_length + 1 < $len) {
        //             break;
        //         }
        //     }
        // }

        // return $str;
    }
}