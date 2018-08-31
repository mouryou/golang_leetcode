func minWindow(s string, t string) string {
    if t == "" {
        return ""
    }
    ls, lt, lc := len(s), len(t), 0
    resL, resR := 0, ls + 1
    var count [256]byte
    for i := 0; i < lt; i++ {
        if count[t[i]] == 0 {
            lc += 1
        }
        count[t[i]] += 1
    }
    l, r := 0, 0
    for r < ls {
        count[s[r]] -= 1
        if count[s[r]] == 0 {
            lc -= 1
        }
        r++
        for lc == 0 {
            count[s[l]] += 1
            if count[s[l]] == 1 {
                lc += 1
                if r - l == lt {
                    return s[l:r]
                }
                if r - l < resR - resL {
                    resR = r
                    resL = l
                }
            }
            l++
        }
    }
    if resR == ls + 1 {
        return ""
    }
    return s[resL:resR]
}