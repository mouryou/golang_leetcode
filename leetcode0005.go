func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// convert "abc" to "^#a#b#c#$"
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

// Manacher's algorithm
func longestPalindrome(s string) string {
    T := preprocess(s)
    n := len(T)
    P := make([]int, n)  // P[i]: the length the longest palindrome whose center is T[i]
    C := 1 // The center of the palindrome having already been found with the rightmost right end
    R := 1 // The right end of this palindrome (inclusive, will be a "#")
    for i := 1; i < n - 1; i += 1 {
        i_mirror := 2 * C - i  // the symmetric position of i with pivot C 
        if R > i {
            // (R - i) / 2 * 2 = R - i is the length of the palindrome whose center is i and right end is R
            // According to the symmetry, if P[i_mirror] < R - i, P[i] will be equal to P[i_mirror]
            // else we initiate P[i] with R - i and expand the border to find the longest palindrome with center i
            P[i] = min(R - i, P[i_mirror])
        } else {
            P[i] = 0
        }
        // Expand the border to find the longest palindrome with center i
        for T[i + 1 + P[i]] == T[i - 1 - P[i]] {
            P[i] += 1
        }
        // Update C and R
        if i + P[i] > R {
            C = i
            R = i + P[i]
        }
    }
    
    // Find the longest palindrome substring
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