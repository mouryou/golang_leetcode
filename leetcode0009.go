func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    origin := x
    reversed := 0
    for x > 0 {
        reversed = reversed * 10 + (x % 10)
        x /= 10
    }
    return origin == reversed
    
}