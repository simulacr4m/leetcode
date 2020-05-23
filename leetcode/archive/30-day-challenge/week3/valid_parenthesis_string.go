func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b
}

func checkValidString(s string) bool {
    lo, hi := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            lo += 1
        } else {
            lo -= 1
        }
        if s[i] != ')' {
            hi += 1
        } else {
            hi -= 1
        }
        if hi < 0 {
            break
        }
        lo = max(lo, 0)
    }
    return lo == 0
}
