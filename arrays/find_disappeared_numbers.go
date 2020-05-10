func findDisappearedNumbers(nums []int) []int {
    var disappeared []int
    var buckets = make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        buckets[nums[i]-1] += 1
    }
    for i := 0; i < len(buckets); i++ {
        if buckets[i] == 0 {
            disappeared = append(disappeared, i+1)
        }
    }
    return disappeared
}
