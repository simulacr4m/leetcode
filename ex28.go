func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    lenNeedle := len(needle)
    for i := 0; i < len(haystack)-lenNeedle+1; i++ {
        if haystack[i:i+lenNeedle] == needle {
            return i
        }
    }
    return -1
}
