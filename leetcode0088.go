func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i := m + n - 1; i >= n; i-- {
        nums1[i] = nums1[i - n]
    }
    i1, i2 := n, 0
    for i := 0; i < m + n; i++ {
        if i1 == m + n {
            nums1[i] = nums2[i2]
            i2++
        } else if i2 == n {
            nums1[i] = nums1[i1]
            i1++
        } else if nums1[i1] < nums2[i2] {
            nums1[i] = nums1[i1]
            i1++
        } else {
            nums1[i] = nums2[i2]
            i2++
        }
    }
}