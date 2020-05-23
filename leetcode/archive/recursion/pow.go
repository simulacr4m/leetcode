func myPowHelper(x float64, n int, cache map[int]float64) float64 {
    if n == 1 {
        return x
    } else if n == 0 {
        return 1.0
    } else if _, ok := cache[n]; ok {
        return cache[n]
    }
    
    a := myPowHelper(x, n / 2, cache)
    b := myPowHelper(x, n - n / 2, cache)
    cache[n] = a*b
    return a*b
}

func myPow(x float64, n int) float64 {
    if x == 1.0 {
        return x
    } else if x == -1.0 {
        if n % 2 == 0 {
            return 1.0
        } else {
            return -1.0
        }
    }
    cache := make(map[int]float64)
    res := myPowHelper(x, int(math.Abs(float64(n))), cache)
    if n < 0 {
        res = 1.0 / res
    }
    return res
}
