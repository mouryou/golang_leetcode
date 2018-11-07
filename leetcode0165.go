func compareVersion(version1 string, version2 string) int {
    n1, n2 := toNums(version1), toNums(version2)
    l1, l2 := len(n1), len(n2)
    i := 0
    for i < l1 && i < l2 {
        if n1[i] > n2[i] {
            return 1
        }
        if n1[i] < n2[i] {
            return -1
        }
        i++
    }
    if l1 > l2 {
        for i < l1 {
            if n1[i] != 0 {
                return 1
            }
            i++
        }
    }
    if l2 > l1 {
        for i < l2 {
            if n2[i] != 0 {
                return -1
            }
            i++
        }
    }
    return 0
}

func toNums(v string) []int {
    res := []int {}
    for _, i := range strings.Split(v, ".") {
        n, _ := strconv.Atoi(i)
        res = append(res, n)
    }
    return res
}