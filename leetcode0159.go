func lengthOfLongestSubstringTwoDistinct(s string) int {
    ls := len(s)
    if ls <= 2 {
        return ls
    }
    letters := make(map[byte]int)
    l, r, res := 0, 0, 0
    for r < ls {
        for r < ls && len(letters) <= 2 {
            letters[s[r]] += 1
            r++
        }
        //fmt.Println(l, r)
        if r - l - 1 > res {
            res = r - l - 1
        }
        for len(letters) > 2 {
            letters[s[l]] -= 1
            if letters[s[l]] == 0 {
                delete(letters, s[l])
            }
            l++
        }
    }
    if r - l > res {
        return r - l
    }
    return res
}