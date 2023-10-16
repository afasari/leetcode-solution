class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function myAtoi($s) {
        $newWord ='';
        foreach(str_split($s) as $letter){
            if(is_numeric((int)$letter) && $letter!='e'){
                $newWord = $newWord.$letter;
                echo $newWord.PHP_EOL;
            } else {
                break;
            }
        }
        $number = (int)$newWord;
        if(abs($number)>=2147483648){
            if($number<0) return -2147483648;
            return 2147483647;
        }
        return $number;
    }
}