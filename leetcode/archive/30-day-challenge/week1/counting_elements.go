func search(arr []int, key int) int {
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        if arr[mid] < key {
            lo = mid+1
        } else if arr[mid] > key {
            hi = mid-1
        } else {
            return key
        }
    }
    return -1
}


func countElements(arr []int) int {
    sort.Ints(arr)
    count := 0
    for i := 0; i < len(arr); i++ {
        x := arr[i]
        index := search(arr[i:], x+1)
        if index >= 0 {
            count += 1
        }
    }
    return count
}
