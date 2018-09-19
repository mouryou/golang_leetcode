import "unicode"

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l, r := 0, len(s) - 1
    for l < r {
        for l < r && !isAlphanumeric(s[l]) {
            l++
        }
        for l < r && !isAlphanumeric(s[r]) {
            r--
        }
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

func isAlphanumeric(b byte) bool {
    return unicode.IsDigit(rune(b)) || unicode.IsLetter(rune(b))
}