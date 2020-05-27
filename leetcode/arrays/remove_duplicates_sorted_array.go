func removeDuplicates(nums []int) int {
    k := 0
    for i := 0; i < len(nums); i++ {
        j := i
        for j < len(nums)-1 && nums[j] == nums[j+1] {
            j += 1
        }
        tmp := nums[j]
        nums[j] = nums[k]
        nums[k] = tmp
        k += 1
        i = j
    }
    return k
}
