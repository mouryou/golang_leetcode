func findMissingRanges(nums []int, lower int, upper int) []string {
    nums = append(nums, upper + 1)
    ranges := map[int][2]int {}
    for _, n := range nums {
        _, ok := ranges[n] 
        if !ok {
            lbs, lok := ranges[n - 1] // left borders
            rbs, rok := ranges[n + 1] // right borders
            lb, rb := n, n
            if lok {
                lb = lbs[0]
            }
            if rok {
                rb = rbs[1]
            }
            r := [2]int {lb, rb}
            ranges[n] = r
            if lok {
                ranges[lb] = r
                ranges[n - 1] = r
            }
            if rok {
                ranges[rb] = r
                ranges[n + 1] = r
            }
        }
    }
    //fmt.Println(ranges)
    sortedRanges := []int {}
    for k, v := range ranges {
        if k == v[0] {
            sortedRanges = append(sortedRanges, k)
        }
    }
    sort.Ints(sortedRanges)
    //fmt.Println(sortedRanges)
    res := []string {}
    low := lower
    for _, b := range sortedRanges {
        if b == low + 1 {
            res = append(res, strconv.Itoa(low))
        } else if b != low {
            res = append(res, strconv.Itoa(low) + "->" + strconv.Itoa(b - 1))
        }
        low = ranges[b][1] + 1
    }
    return res
}