import "math"

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    //do binary search in the shorter array
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    tl := len(nums1) + len(nums2) // total length
    i, j := 0, len(nums1) + 1 //i is the length of the left part of nums1
    for i < j {
        m := (i + j) / 2 //mid in nums1
        m2 := tl / 2 - m  // mid in nums2
        var l1, r1, l2, r2 int
        if m == 0 {
            l1 = math.MinInt64
        } else {
            l1 = nums1[m - 1]
        }
        if m == len(nums1) {
            r1 = math.MaxInt64
        } else {
            r1 = nums1[m]
        }
        if m2 == 0 {
            l2 = math.MinInt64
        } else {
            l2 = nums2[m2 - 1]
        }
        if m2 == len(nums2) {
            r2 = math.MaxInt64
        } else {
            r2 = nums2[m2]
        }
        if l1 > r2 {
            j = m
        } else if r1 < l2 {
            i = m + 1
        } else if tl % 2 == 1 {
            return float64(min(r1, r2))
        } else {
            return float64(max(l1, l2) + min(r1, r2)) / 2.0
        }
    }
    return 0

}