function twoSum(nums: number[], target: number): number[] {
    let hashMap: Map<number, number> = new Map();
    let length: number = nums.length;
    for (let i=0; i < length; i++){
        if (hashMap[nums[i]] !== undefined) {
            return [hashMap[nums[i]], i]
        }
        hashMap[target-nums[i]] = i
    }

    return []
    
};