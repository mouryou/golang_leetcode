func lengthOfLastWord(s string) int {
    result := 0
    start := false
    i := len(s) - 1
    for i >= 0 {
        if start {
            if s[i] == ' ' {
                return result
            } else {
                result += 1
            }
        } else {
            if s[i] != ' ' {
                start = true
                result = 1
            }
        }
        i--
    }
    return result
}