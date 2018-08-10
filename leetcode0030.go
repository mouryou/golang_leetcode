func findSubstring(s string, words []string) []int {
    wc := len(words) // word counts
    res := make([]int, 0)
    if wc == 0 {
        return res
    }
    wl := len(words[0]) // word length
    sl := len(s) // s length
    owcm := make(map[string]int)  // original word counts map
    for _, w := range words {
        if _, ok := owcm[w]; ok {
            owcm[w] += 1
        } else {
            owcm[w] = 1
        }
    }
    for i := 0; i < wl && i <= sl - wl * wc; i++ {
        wcm := make(map[string]int)
        for k, v := range owcm {
            wcm[k] = v
        }
        lb := i // left border
        for j := i; j <= sl - wl; j += wl {
            cw := s[j:j + wl] // current word
            if _, ok := wcm[cw]; ok {
                wcm[cw] -= 1
                for wcm[cw] < 0 {
                    wcm[s[lb:lb + wl]] += 1
                    lb += wl
                }
                if j == lb + (wc - 1) * wl {
                    res = append(res, lb)
                }
            } else {
                for ; lb < j; lb += wl {
                    wcm[s[lb:lb + wl]] += 1
                }
                lb = j + wl
            }
            if lb > sl - wl * wc {
                break
            }
        }
    } 
    return res
}