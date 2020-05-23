package main

import "fmt"

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for k := i+1; k < len(nums); k++ {
            if nums[k] == target - nums[i] {
                return []int{i, k}
            }
        }
    }
    return []int{-1, -1}
}

/* Another Solution */

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
    cpy := make(map[int][]int)
    for i := 0; i < len(nums); i++ {
        if cpy[nums[i]] != nil {
            cpy[nums[i]] = append(cpy[nums[i]], i)
        } else {
            cpy[nums[i]] = []int{i}
        }
    }
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        x := binarySearch(nums, target - nums[i], i+1)
        if x >= 0 {
            if nums[x] == nums[i] {
                return []int{cpy[nums[x]][0], cpy[nums[x]][1]}
            }
            return []int{cpy[nums[x]][0], cpy[nums[i]][0]}
        }
    }
    return nil
}
