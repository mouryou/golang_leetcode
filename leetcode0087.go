func isScramble(s1 string, s2 string) bool {
    dp := make(map[[2]string]bool)     
    return search(s1, s2, dp)
}

func search(s1, s2 string, dp map[[2]string]bool) bool {
    l := len(s1)
    if l == 1 {
        return s1 == s2
    }
    count := make(map[byte]int)
    for i := 0; i < l; i++ {
        count[s1[i]] += 1
    }
    for i := 0; i < l; i++ {
        if count[s2[i]] == 0 {
            return false
        }
        count[s2[i]] -= 1
    }
    key := [2]string {s1, s2}
    if res, ok := dp[key]; ok {
        return res
    }
    for i := 1; i < l; i++ {
        if (search(s1[:i], s2[:i], dp) && search(s1[i:], s2[i:], dp)) || (search(s1[:i], s2[l - i:], dp) && search(s1[i:], s2[:l - i], dp)) {
            dp[key] = true
            return true
        }
    }
    dp[key] = false
    return false
}
