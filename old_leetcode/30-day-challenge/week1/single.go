func singleNumber(nums []int) int {
    
    count := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        count[nums[i]] += 1
    }
    for key, val := range count {
        if val == 1 {
            return key;
        }
    }
    return -1;
}
