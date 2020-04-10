func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    i, j, k, size := 0, 0, 0, len(nums1) + len(nums2)
    nums3 := make([]int, size)
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            nums3[k] = nums1[i]
            i += 1
        } else {
            nums3[k] = nums2[j]
            j += 1
        }
        k += 1
    }
    for i < len(nums1) {
        nums3[k] = nums1[i]
        i += 1
        k += 1
    }
    for j < len(nums2) {
        nums3[k] = nums2[j]
        j += 1
        k += 1
    }
    median := 0.0
    if size % 2 == 0 {
        median = float64(nums3[size/2] + nums3[size/2 - 1]) / 2.0
    } else {
        median = float64(nums3[size/2])
    }
    return median
}
