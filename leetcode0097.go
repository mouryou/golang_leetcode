func isInterleave(s1 string, s2 string, s3 string) bool {
    return search(s1, s2, s3, 0, 0, make(map[[2]int]bool))
}

func search(s1, s2, s3 string, i1, i2 int, dp map[[2]int]bool) bool {
    key := [2]int {i1, i2}
    res, ok := dp[key]
    if ok {
        return res
    }
    l1, l2 := len(s1), len(s2)
    for i1 < l1 && i2 < l2 {
        if s1[i1] == s3[i1 + i2] && s2[i2] == s3[i1 + i2] {
            res = search(s1, s2, s3, i1, i2 + 1, dp) || search(s1, s2, s3, i1 + 1, i2, dp)
            dp[key] = res
            return res
        } else if s1[i1] == s3[i1 + i2] {
            i1++
        } else if s2[i2] == s3[i1 + i2] {
            i2++
        } else {
            dp[key] = false
            return false
        }
    }
    if i1 == l1 {
        res = s2[i2:] == s3[i1 + i2:]
    }
    if i2 == l2 {
        res = s1[i1:] == s3[i1 + i2:]
    }
    dp[key] = res
    return res
}