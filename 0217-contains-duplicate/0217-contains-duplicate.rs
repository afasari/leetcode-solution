use std::collections::HashMap;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut hashMap = HashMap::new();

        for n in nums{
            if (hashMap.contains_key(&n)){
                return true;
            }
            hashMap.insert(n, true);
        }

        return false;
    }
}

    // contacts.insert("Robert", "956-1745");

    // // Takes a reference and returns Option<&V>
    // match contacts.get(&"Daniel") {
    //     Some(&number) => println!("Calling Daniel: {}", call(number)),
    //     _ => println!("Don't have Daniel's number."),
    // }