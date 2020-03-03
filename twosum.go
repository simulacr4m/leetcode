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
