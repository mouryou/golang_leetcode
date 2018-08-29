func plusOne(digits []int) []int {
    c := 1
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i] += c
        c = digits[i] / 10
        if c == 0 {
            return digits
        }
        digits[i] %= 10
    }
    if c == 1 {
        return append([]int {1}, digits...)
    }
    return digits
}