func isValid(s string) bool {
        // 1. Early exit for odd length
    if len(s)%2 != 0 {
        return false
    }

    // 2. Map closing brackets to opening ones
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    // 3. Stack to store opening brackets
    stack := []rune{}

    for _, char := range s {
        // If it's a closing bracket
        if opener, ok := pairs[char]; ok {
            // Check if stack is empty or top doesn't match
            if len(stack) == 0 || stack[len(stack)-1] != opener {
                return false
            }
            // Pop from stack
            stack = stack[:len(stack)-1]
        } else {
            // Push opening bracket onto stack
            stack = append(stack, char)
        }
    }

    // 4. Return true if all brackets were matched
    return len(stack) == 0
}