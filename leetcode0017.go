func letterCombinations(digits string) []string {
    res := make([]string, 0)
    if len(digits) == 0 {
        return res
    }
    dic := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    var t []string
    if len(digits) == 1 {
        t = []string {""}
    } else {
        t = letterCombinations(digits[1:])
    }
    for _, b := range dic[digits[0]] {
        for _, c := range t {
            res = append(res, string(b) + c)
        }
    }
    return res
    
}