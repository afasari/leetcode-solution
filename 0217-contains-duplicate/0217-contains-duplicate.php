class Solution {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function containsDuplicate($nums) {
        $hashMap = array();

        for ($i = 0; $i < count($nums); $i++){
            if ($hashMap[$nums[$i]]){
                return true;
            }
            $hashMap[$nums[$i]] = true;
        }

        return false;
    }
}