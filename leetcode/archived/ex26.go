func removeDuplicates(nums []int) int {
    k, r := 0, 0
    for i := 1; i < len(nums); i++ {
        for k < len(nums)-1 && nums[k] == nums[k+1] {
            k += 1
        }
        if k < len(nums)-1 {
            nums[i] = nums[k+1]
            r = i
            k += 1
        }
    }
    return r+1
}
