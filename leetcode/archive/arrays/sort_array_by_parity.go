func sortArrayByParity(A []int) []int {
    if len(A) == 0 {
        return A
    }
    i := 0;
    for j := 0; j < len(A); j++ {
        if A[j] % 2 == 0 {
            tmp := A[i]
            A[i] = A[j]
            A[j] = tmp
            i += 1
        }
    }
    return A
}
