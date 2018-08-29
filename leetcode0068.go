func fullJustify(words []string, maxWidth int) []string {
    result := make([]string, 0)
    startI := 0
    curLen := len(words[0])
    for i := 1; i < len(words); i++ {
        if curLen + len(words[i]) + i - startI - 1 >= maxWidth {
            spaceNum := i - startI - 1
            if spaceNum == 0 {
                result = append(result, words[startI] + strings.Repeat(" ", maxWidth - curLen))
            } else {
                spaceLen := (maxWidth - curLen) / spaceNum
                longSpaceNum := (maxWidth - curLen) % spaceNum
                s := ""
                for j := 0; j < longSpaceNum; j++ {
                    s += words[startI + j] + strings.Repeat(" ", spaceLen + 1)
                }
                for j := longSpaceNum; j < spaceNum; j++ {
                    s += words[startI + j] + strings.Repeat(" ", spaceLen)
                }
                result = append(result, s + words[i - 1])
            }
            startI = i
            curLen = len(words[i])
        } else {
            curLen += len(words[i])
        }
        
    }
    s := strings.Join(words[startI:], " ")
    return append(result, s + strings.Repeat(" ", maxWidth - len(s)))
}