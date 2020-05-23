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

/* More efficient solution */

func findDisappearedNumbers(nums []int) []int {
    var disappeared []int
    for i := 0; i < len(nums); i++ {
        for nums[i] != 0 && nums[nums[i]-1] != 0 {
            index := nums[i]-1
            nums[i] = nums[index]
            nums[index] = 0
        }
    }
    
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            disappeared = append(disappeared, i+1)
        }
    }
    return disappeared
}
