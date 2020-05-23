func subarraySum(nums []int, k int) int {
    count := 0
    sum := make([]int, len(nums)+1)
    sum[0] = 0
    for i := 1; i <= len(nums); i++ {
        sum[i] = sum[i-1] + nums[i-1]
    }
    for start := 0; start < len(nums); start++ {
        for end := start+1; end <= len(nums); end++ {
            if sum[end] - sum[start] == k {
                count += 1
            }
        }
    }
    return count
}
