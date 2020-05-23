func fibHelper(N int, cache []int) {
    if N < 2 {
        cache[N] = N
        return
    }
    fibHelper(N-2, cache)
    fibHelper(N-1, cache)
    cache[N] = cache[N-1] + cache[N-2]
}

func fib(N int) int {
    cache := make([]int, N+1)
    fibHelper(N, cache)
    return cache[N]
}
