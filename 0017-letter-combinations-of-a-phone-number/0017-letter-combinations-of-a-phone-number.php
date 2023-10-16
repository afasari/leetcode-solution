class Solution {

    /**
     * @param String $digits
     * @return String[]
     */
    function letterCombinations($digits) {
        $map = [
            2 => ['a', 'b', 'c'],
            3 => ['d', 'e', 'f'],
            4 => ['g', 'h', 'i'],
            5 => ['j', 'k', 'l'],
            6 => ['m', 'n', 'o'],
            7 => ['p', 'q', 'r', 's'],
            8 => ['t', 'u', 'v'],
            9 => ['w', 'x', 'y', 'z']
        ];
        $digitArr = str_split($digits);
		if(strlen($digits) == 0) {
			return [];
		}
		
		if(strlen($digits) == 1) {
			return $map[$digitArr[0]];
		}
		if(strlen($digits) == 2) {
			for($i = 0 ;$i < count($map[$digitArr[0]]); $i++) {
				for($j = 0; $j < count($map[$digitArr[1]]); $j++) {
					$results[] = $map[$digitArr[0]][$i] . $map[$digitArr[1]][$j];
				}
			}
			return $results;
		}
		
		// if count > 2 go recursion
			for($i = 0 ;$i < count($map[$digitArr[0]]); $i++) {
				$res = $this->letterCombinations(substr($digits, 1));
				foreach($res as $r){
					$results[] = $map[$digitArr[0]][$i] . $r;
				}

			}
			return $results;
    }
}