func sortedSquares(A []int) []int {
    squares := make([]int, len(A))
    for i := 0; i < len(squares); i++ {
        squares[i] = A[i] * A[i]
    }
    sort.Ints(squares)
    return squares
}
