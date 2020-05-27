func moveZeroes(nums []int)  {
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            k := i+1
            for ; k < len(nums); k++ {
                if nums[k] != 0 {
                    break
                }
            }
            if k >= len(nums) {
                break
            }
            tmp := nums[i]
            nums[i] = nums[k]
            nums[k] = tmp
        }
    }
}
