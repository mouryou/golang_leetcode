func isOneEditDistance(s string, t string) bool {
    ls, lt := len(s), len(t)
    if ls < lt {
        s, t, ls, lt = t, s, lt, ls
    }
    if ls - lt == 1 {
        i := 0
        for ; i < lt; i++ {
            if s[i] != t[i] {
                break
            }
        }
        for ; i < lt; i++ {
            if s[i + 1] != t[i] {
                break
            }
        }
        return i == lt
    }
    if ls == lt {
        diff := 0
        for i := 0; i < ls; i++ {
            if s[i] != t[i] {
                diff += 1
                if diff > 1 {
                    return false
                }
            }
        }
        return diff == 1
    }
    return false
}