func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    results := [][]string {}
    wordSet := make(map[string]bool)
    for _, w := range wordList {
        wordSet[w] = true
    }
    delete(wordSet, beginWord)
    if !wordSet[endWord] {
        return results
    }
    prevs := make(map[string][]string)
    border := []string {beginWord}
    
    for len(wordSet) > 0 {
        nextBorder := []string {}
        for w, _ := range wordSet {
            taken := false
            for _, pw := range border {
                if isAdjacent(w, pw) {
                    if !taken {
                        nextBorder = append(nextBorder, w)
                        taken = true
                    }
                    prevs[w] = append(prevs[w], pw)
                }
            }
        }
        if len(prevs[endWord]) > 0 {
            break
        }
        if len(nextBorder) == 0 {
            return results
        }
        for _, w := range nextBorder {
            delete(wordSet, w)
        }
        border = nextBorder
    }
    if len(prevs[endWord]) == 0 {
        return results
    }
    appendResult(prevs, []string {}, endWord, &results)
    return results
}

func isAdjacent(w1, w2 string) bool {
    l := len(w1)
    d := 0
    for i := 0; i < l; i++ {
        if w1[i] != w2[i] {
            d += 1
        }
        if d > 1 {
            return false
        }
    }
    return d == 1
}

func appendResult(prevs map[string][]string, curResult []string, curWord string, results *[][]string) {
    if len(prevs[curWord]) == 0 {
        *results = append(*results, append([]string {curWord}, curResult...))
    } else {
        for _, w := range prevs[curWord] {
            appendResult(prevs, append([]string {curWord}, curResult...), w, results)
        }
    }
}