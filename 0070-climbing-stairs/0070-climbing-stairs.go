func climbStairs(n int) int {
    return hackersApproach(n)
}

// O(n) time and O(1) space
func hackersApproach(n int) int{
    next, secondNext := 0, 1
    for ; n > 0; n-- {
        next, secondNext = secondNext, next + secondNext
    }
    return secondNext
}

// O(2^n) time and O(n) space
func fibonacci(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    }
    
    return climbStairs(n-1) + climbStairs(n-2)    
}

// O(n) time and O(n) space
var mem map[int]int = map[int]int{}

func dpMemoization(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    } else if val, ok := mem[n]; ok {
        return val
    }

    res := climbStairs(n-1) + climbStairs(n-2)
    mem[n] = res
    return res
}

// O(n) time and O(1) space
func iterative(n int) int {
    res := 0
    
    secondNext := 0
    next := 0
    for i := 1; i <= n; i++ {
        if i == 1 {
            res = 1
        } else if i == 2 {
            res = 2
        } else {
            res = secondNext + next
        }
        
        next = secondNext
        secondNext = res
    }
    
    return res
}