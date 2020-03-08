func merge(nums1 []int, m int, nums2 []int, n int)  {
    k := 0
    for i := m; i < m+n && k < n; i++ {
        nums1[i] = nums2[k]
        k += 1
    }
    sort.Ints(nums1)
}
