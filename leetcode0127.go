func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := map[string]bool {}
    for _, word := range wordList {
        wordSet[word] = true
    }
    if !wordSet[endWord]{
        return 0
    }
    visited := map[string]bool {}

    beginSet := map[string]bool {}
    endSet := map[string]bool {}
    beginSet[beginWord] = true
    endSet[endWord] = true

    length := 2
    for len(beginSet) > 0 && len(endSet) > 0 {
        if len(beginSet) > len(endSet) {
            beginSet, endSet = endSet, beginSet
        }
        
        next := map[string]bool {}
        for word := range beginSet {
            visited[word] = true
            
            wordBytes := []byte(word)
            for i := 0; i < len(wordBytes); i++ {
                for ch := byte('a'); ch <= byte('z'); ch++ {
                    if ch == wordBytes[i] {
                        continue
                    }
                    wordBytes[i] = ch
                    neigh := string(wordBytes)
                    if !wordSet[neigh] {
                        continue
                    }
                    if endSet[neigh] {
                        return length
                    }
                    if visited[neigh] {
                        continue
                    }
                    next[neigh] = true            
                }
                wordBytes[i] = word[i]
            }
        }

        beginSet = next
        length++
    }
    
    return 0
}