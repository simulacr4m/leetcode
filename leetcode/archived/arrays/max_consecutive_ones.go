func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    for i := 0; i < len(nums); i++ {
        j := i
        for ; j < len(nums); j++ {
            if nums[j] == 0 {
                break
            }
        }
        if max < j-i {
            max = j-i
        }
    }
    return max
}
