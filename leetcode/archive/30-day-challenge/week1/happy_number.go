func isHappy(n int) bool {
    for n != 1 {
        sum := 0
        for n != 0 {
            sum += (n % 10) * (n % 10)
            n /= 10
        }
        n = sum
        if n < 7 {
            break
        }
    }
    return n == 1
}
