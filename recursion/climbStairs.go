
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    cache := make([]int, n)
    cache[0] = 1
    cache[1] = 2
    
    for i := 2; i < n; i++ {
        cache[i] = cache[i-2] + cache[i-1]
    }
    return cache[n-1]
}
