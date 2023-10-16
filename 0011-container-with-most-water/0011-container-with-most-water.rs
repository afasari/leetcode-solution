impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut max :i32 = 0;
        let lh = height.len();
        if lh < 2{
            return 0;
        }

        let mut left :usize = 0;
        let mut right :usize = lh-1;
        while left < right{
            let mut area :i32 = 0;
            if height[left] > height[right]{
                area = height[right] * (right-left) as i32;
                right -= 1;
            } else{
                area = height[left] * (right-left) as i32;
                left += 1;
            }

            if area > max{
                max =area;
            }
        }
        max
    }
}
