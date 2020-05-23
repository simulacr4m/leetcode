func productExceptSelf(nums []int) []int {
    output := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        product := 1
        for k := 0; k < len(nums); k++ {
            if k != i {
                product *= nums[k]
            }
        }
        output[i] = product
    }
    return output
}
