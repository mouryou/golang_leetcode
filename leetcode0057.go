/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func insert(intervals []Interval, newInterval Interval) []Interval {
    li := len(intervals)
    if li == 0 {
        return []Interval { newInterval }
    }
    if newInterval.Start > intervals[li - 1].End {
        return append(intervals, newInterval)
    }
    iStart, iEnd := search(intervals, newInterval.Start), search(intervals, newInterval.End)
    var l, r int
    l = iStart / 2
    if iStart % 2 == 1 {
        newInterval.Start = intervals[iStart / 2].Start
    }
    if iEnd == li * 2 {
        r = li
    } else if iEnd % 2 == 1 || newInterval.End == intervals[iEnd / 2].Start {
        newInterval.End = intervals[iEnd / 2].End
        r = iEnd / 2 + 1
    } else {
        r = iEnd / 2
    }
    return append(intervals[:l], append([]Interval {newInterval}, intervals[r:]...)...)
}

func search(intervals []Interval, target int) int {
    li := len(intervals)
    if target <= intervals[0].Start {
        return 0
    }
    if target > intervals[li - 1].End {
        return 2 * li
    }
    l, r := 0, 2 * li
    for l < r {
        m := (l + r) / 2
        if getValue(intervals, m) > target {
            r = m
        } else if getValue(intervals, m) < target {
            l = m + 1
        } else {
            return m
        }
    }    
    return l
}
func getValue(intervals []Interval, i int) int {
    if i % 2 == 0 {
        return intervals[i / 2].Start
    } else {
        return intervals[i / 2].End
    }
}