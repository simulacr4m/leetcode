func search(nums []int, val, start int) int {
    lo, hi := start, len(nums)-1
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        if nums[mid] < val {
            lo = mid + 1
        } else if nums[mid] > val {
            hi = mid - 1
        } else {
            return mid
        }
    }
    return -1
}

func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    fmt.Println(nums)
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for k := i+1; k < len(nums); k++ {
            if k != i+1 && nums[k] == nums[k-1] {
                continue
            }
            d := search(nums, -(nums[i] + nums[k]), k+1)
            if d > 0 {
                res = append(res, []int{nums[i], nums[k], nums[d]})
            }
        }
    }
    return res
}
