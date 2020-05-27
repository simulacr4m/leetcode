func binSearch(arr []int, target int) int {
    i, j := 0, len(arr)-1
    for i <= j {
        mid := (i + j) / 2
        if arr[mid] < target {
            i = mid+1
        } else if arr[mid] > target {
            j = mid-1
        } else {
            return mid
        }
    }
    return -1
}

func checkIfExist(arr []int) bool {
    sort.Ints(arr)
    for i := 0; i < len(arr); i++ {
        index := binSearch(arr, 2 * arr[i])
        if index >= 0 && index != i {
            return true
        }
    }
    return false
}
