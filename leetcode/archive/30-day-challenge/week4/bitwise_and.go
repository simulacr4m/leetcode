func rangeBitwiseAnd(m int, n int) int {
    for i := m+1; i <= n; i++ {
        m &= i
    }
    return m
}
