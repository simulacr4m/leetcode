func duplicateZeros(arr []int)  {
    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            for j := len(arr)-1; j >= i+1; j-- {
                arr[j] = arr[j-1]
            }
            i += 1
        }
    }
}