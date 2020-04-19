func findNumbers(nums []int) int {
    even := 0
    for i := 0; i < len(nums); i++ {
        cpy, digits := nums[i], 0
        for cpy != 0 {
            cpy /= 10
            digits += 1
        }
        if digits % 2 == 0 {
            even += 1
        }
    }
    return even
}
