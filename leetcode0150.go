func evalRPN(tokens []string) int {
    stack := []int {}
    for _, t := range tokens {
        l := len(stack)
        switch t {
        case "+":
            stack[l - 2] += stack[l - 1]
            stack = stack[:l - 1]
        case "-":
            stack[l - 2] -= stack[l - 1]
            stack = stack[:l - 1]
        case "*":
            stack[l - 2] *= stack[l - 1]
            stack = stack[:l - 1]
        case "/":
            stack[l - 2] /= stack[l - 1]
            stack = stack[:l - 1]
        default:
            n, _ := strconv.Atoi(t)
            stack = append(stack, n)
        }
    }
    return stack[0]
}