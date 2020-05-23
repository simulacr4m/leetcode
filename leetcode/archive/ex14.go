func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    min := strs[0]
    for i := 0; i < len(strs); i++ {
        if len(strs[i]) < len(min) {
            min = strs[i]
        }
    }
    prefix := ""
    for i := 1; i <= len(min); i++ {
        prefix = strs[0][0:i]
        for k := 1; k < len(strs); k++ {
            if strs[k][0:i] != prefix {
                return strs[0][0:i-1]
            }
        }
    }
    return prefix
}
