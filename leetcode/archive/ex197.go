func hasAlternatingBits(n int) bool {
    var prev int
    if n & 1 == 0 {
        prev = 1
    } else {
        prev = 0
    }
    for n != 0 {
        r := n & 1
        if (r == 0 && prev != 1) || (r == 1 && prev != 0) {
            return false
        }
        prev = r
        n >>= 1
    }
    return true
}
