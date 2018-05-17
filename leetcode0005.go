func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func preprocess(s string) string {
    n := len(s)
    if n == 0 {
        return "^$"
    }
    ret := "^"
    for _, r := range s {
        ret += "#" + string(r)
    }
    ret += "#$"
    return ret
}


func longestPalindrome(s string) string {
    T := preprocess(s)
    n := len(T)
    P := make([]int, n)
    C, R := 0, 0
    for i := 1; i < n - 1; i += 1 {
        i_mirror := 2 * C - i
        if R > i {
            P[i] = min(R - i, P[i_mirror])
        } else {
            P[i] = 0
        }
        for T[i + 1 + P[i]] == T[i - 1 - P[i]] {
            P[i] += 1
        }
        if i + P[i] > R {
            C = i
            R = i + P[i]
        }
    }
    
    maxLen, centerIndex := 0, 0
    for i := 1; i < n - 1; i += 1 {
        if P[i] > maxLen {
            maxLen = P[i]
            centerIndex = i
        }
    }
    start := (centerIndex - 1 - maxLen) / 2
    return s[start:(start + maxLen)]
}