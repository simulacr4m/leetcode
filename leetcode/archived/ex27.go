func removeElement(nums []int, val int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            for k := i+1; k < len(nums); k++ {
                if nums[k] != val {
                    tmp := nums[i]
                    nums[i] = nums[k]
                    nums[k] = tmp
                }
            }
        }
    }
    r := 0
    for ; r < len(nums); r++ {
        if nums[r] == val {
            break
        }
    }
    return r
}
