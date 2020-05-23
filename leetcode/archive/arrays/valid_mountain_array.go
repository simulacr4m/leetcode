func validMountainArray(A []int) bool {
    if len(A) < 3 {
        return false
    }
    
    j, k := 0, 0
    for i := 0; i < len(A)-1; i++ {
        if A[i] == A[i+1] {
            return false
        }
        if A[i] > A[i+1] {
            j = i
            break
        }
    }
    
    for i := len(A)-1; i > 0; i-- {
        if A[i] == A[i-1] {
            return false
        }
        if A[i] > A[i-1] {
            k = i
            break
        }
    }
    if j+1 >= len(A) || j-1 < 0 || k+1 >= len(A) || k-1 < 0 {
        return false
    }
    return j == k
}
