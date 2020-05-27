func removeElement(nums []int, val int) int {
    last := len(nums)-1
    for i := 0; i <= last; i++ {
        for last >= 0 && nums[last] == val {
            last -= 1
        }
        if i <= last && nums[i] == val {
            tmp := nums[i]
            nums[i] = nums[last]
            nums[last] = tmp
            last -= 1
        }
    }
    return last+1
}
