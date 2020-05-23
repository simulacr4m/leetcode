func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func thirdMax(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    } else if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    max := math.MinInt32
    secondMax := math.MinInt32
    thirdMax := math.MinInt32
    for i := 0; i < len(nums); i++ {
        if max < nums[i] {
            max = nums[i]
        }
    }
    
    for i := 0; i < len(nums); i++ {
        if nums[i] < max && secondMax < nums[i] {
            secondMax = nums[i]
        }
    }
    
    minPresent := false
    for i := 0; i < len(nums); i++ {
        if nums[i] < secondMax && thirdMax < nums[i] {
            thirdMax = nums[i]
        } else if nums[i] == math.MinInt32 {
            minPresent = true
        }
    }
    if thirdMax == secondMax || (thirdMax == math.MinInt32 && !minPresent) {
        thirdMax = max
    } 
    return thirdMax
}
