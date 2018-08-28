import "fmt"

func groupAnagrams(strs []string) [][]string {
    res := make([][]string, 0)
    count := make(map[string][]string)
    for _, s := range strs {
        k := getKey(s)
        count[k] = append(count[k], s)
    }
    for _, v := range count {
        res = append(res, v)
    }
    return res
}

func getKey(s string) string {
    count := make(map[rune]int)
    for _, r := range s {
        count[r] = count[r] + 1
    }
    res := ""
    for _, r := range "abcdefghijklmnopqrstuvwxyz" {
        if count[r] > 0 {
            res += string(r) + string(count[r])
        }
    }
    return res
}