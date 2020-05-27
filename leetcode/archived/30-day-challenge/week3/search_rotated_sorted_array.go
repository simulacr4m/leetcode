func searchArray(nums []int, target, lo, hi int) int {
    for lo <= hi {
        mid := (lo + hi) / 2
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


func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    pivotPoint := 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            pivotPoint = i-1
            break
        }
    }
    fmt.Println(pivotPoint)
    index := searchArray(nums, target, 0, pivotPoint)
    if index == -1 {
        index = searchArray(nums, target, pivotPoint+1, len(nums)-1)
    }
    return index
}
