func lengthOfLongestSubstring(s string) int {
    counts := make(map[byte]int) //number of every unique character
    res := 0
    //i, j: left (inclusive) and right (exclusive) end of current substring
    for i, j := 0, 0; j < len(s); {
        //add the next character into the current string
        r := s[j] 
        j += 1
        count, ok := counts[r]
        if ok {
            counts[r] += 1
        } else {
            counts[r] = 1
        }
        //remove the first character of the current string until there is no repeating character
        for len(counts) != j - i {
            r = s[i]
            count = counts[r]
            if count == 1 {
                delete(counts, r)
            } else {
                counts[r] -= 1
            }
            i += 1
        }
        //update the result
        if j - i > res {
            res = j - i
        }
    }
    return res
}