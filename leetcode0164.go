func maximumGap(nums []int) int {
    set := make(map[int]bool)
    for _, n := range(nums) {
        set[n] = true
    }
    
    l := len(set)
    if l <= 1 {
        return 0
    }
    maxn, minn := nums[0], nums[0]
    for n, _ := range set {
        if n > maxn {
            maxn = n
        }
        if n < minn {
            minn = n
        }
    }
    breadth := (maxn - minn) / (l - 1)
    nb := (maxn - minn) / breadth + 1
    buckets := make([][2]int, nb)
    for i := 0; i < nb; i++ {
        buckets[i] = [2]int {-1, -1}
    }
    for n, _  := range set {
        i := (n - minn) / breadth
        if n > buckets[i][1] {
            buckets[i][1] = n
        }
        if n < buckets[i][0] || buckets[i][0] == -1 {
            buckets[i][0] = n
        }
    }
    res := 0
    lastMax := buckets[0][1]
    for i := 1; i < nb; i++ {
        if buckets[i][0] != -1 {
            if buckets[i][0] - lastMax > res {
                res = buckets[i][0] - lastMax
            }
            lastMax = buckets[i][1]
        }
    }
    return res
}