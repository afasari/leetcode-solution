use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = HashMap::new();
        let mut ret: Vec<i32> = Vec::new();

        for i in 0..nums.len(){
            match m.get(&(target - nums[i])) {
                Some(idx) => {
                    ret.push(idx-0);
                    ret.push(i as i32);
                    return ret;
                }
                None => {
                    m.insert(&nums[i], i as i32);
                }
            }
        }
        return ret;
    }
}
