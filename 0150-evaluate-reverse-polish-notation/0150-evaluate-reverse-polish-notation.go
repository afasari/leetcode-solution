func evalRPN(tokens []string) int {
        stack := []int{}

    for _, t := range tokens {
        // 1. Check if token is an operator
        if t == "+" || t == "-" || t == "*" || t == "/" {
            // 2. Pop the two most recent operands
            // Important: second pop is the left-side operand
            v2 := stack[len(stack)-1]
            v1 := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            // 3. Apply operation and push result
            switch t {
            case "+":
                stack = append(stack, v1+v2)
            case "-":
                stack = append(stack, v1-v2)
            case "*":
                stack = append(stack, v1*v2)
            case "/":
                stack = append(stack, v1/v2)
            }
        } else {
            // 4. Token is a number, convert and push
            val, _ := strconv.Atoi(t)
            stack = append(stack, val)
        }
    }

    return stack[0]
}