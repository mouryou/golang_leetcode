func canCompleteCircuit(gas []int, cost []int) int {
    start, cum, all := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        if d := gas[i] - cost[i]; cum + d < 0 {
            start, cum, all = i + 1, 0, all + d
        } else {
            cum, all = cum + d, all + d
        }
    }
    if all >= 0 {
        return start
    }
    return -1
}