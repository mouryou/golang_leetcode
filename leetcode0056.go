/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    if len(intervals) <= 1 {
        return intervals
    }
    sort.Slice(intervals, func (i, j int) bool { return intervals[i].Start < intervals[j].Start })
    curInterval := intervals[0]
    result := make([]Interval, 0)
    for _, interval := range intervals {
        if interval.Start <= curInterval.End {
            if interval.End > curInterval.End {
                curInterval.End = interval.End
            }
        } else {
            result = append(result, curInterval)
            curInterval = interval
        }
    }
    result = append(result, curInterval)
    return result
}