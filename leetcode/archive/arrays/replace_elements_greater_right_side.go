func replaceElements(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        max := math.MinInt32
        for k := i+1; k < len(arr); k++ {
            if max < arr[k] {
                max = arr[k]
            }
        }
        arr[i] = max
    }
    arr[len(arr)-1] = -1
    return arr
}
