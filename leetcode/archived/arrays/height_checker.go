func radixSort(heights []int) int {
    toMove := 0
    buckets := make([]int, 100)
    for i := 0; i < len(heights); i++ {
        buckets[heights[i]-1] += 1
    }
    
    i := 0
    for j := 0; j < len(heights); j++ {
        for buckets[i] == 0 {
            i += 1
        }
        if heights[j] != i+1 {
            toMove += 1
        }
        heights[j] = i+1
        buckets[i] -= 1
    }
    return toMove
}


func heightChecker(heights []int) int {
    return radixSort(heights)
}
