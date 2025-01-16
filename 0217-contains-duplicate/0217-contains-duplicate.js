/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    let hashmap = new Map();
    for (let i=0;i < nums.length; i++) {
        if (hashmap.has(nums[i])) {
            return true
        }

        
        hashmap.set(nums[i], true);
    }
    return false
};

// // Create a new empty hashmap
// let hashmap: Map<KeyType, ValueType> = new Map();

// // Add key-value pairs to the hashmap
// hashmap.set(key1, value1);
// hashmap.set(key2, value2);

// // Get the value associated with a key
// let value = hashmap.get(key);

// // Check if a key exists in the hashmap
// let exists = hashmap.has(key);

// // Delete a key-value pair from the hashmap
// hashmap.delete(key);

// // Total number of key-value pair hashmap contains.
// let total = hashMap.size;

// //Removes the all key-value pairs from the map() object
// hashMap.clear();