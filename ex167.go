func binarySearch(nums []int, target int, lo int) int {
    hi := len(nums)-1
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        if nums[mid] < target {
            lo = mid+1
        } else if nums[mid] > target {
            hi = mid-1
        } else {
            return mid
        }
    }
    return -1
}

func twoSum(nums []int, target int) []int {
    
    for i := 0; i < len(nums); i++ {
        x := binarySearch(nums, target - nums[i], i+1)
        if x >= 0 {
            return []int{i+1, x+1}
        }
    }
    return nil
}
