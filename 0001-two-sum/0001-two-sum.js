/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let hashmap = new Map();
    for (let i=0;i < nums.length; i++) {
        if (hashmap.has(nums[i])) {
            return [hashmap.get(nums[i]), i]
        }
        
        hashmap.set(target-nums[i], i);
    }
    return []
};