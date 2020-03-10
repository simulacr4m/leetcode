func repeatedSubstringPattern(s string) bool {
    for i := 1; i <= len(s) / 2; i++ {
        repeat := s[0:i]
        if len(s) % i != 0 {
            continue
        }
        repeats := true
        for k := 0; k < len(s); k += i {
            if repeat != s[k:k+i] {
                repeats = false
                break
            }
        }
        if repeats {
            return true
        }
    }
    return false
}
