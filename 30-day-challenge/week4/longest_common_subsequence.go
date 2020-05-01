func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func longestHelper(text1, text2 string, m, n int) int {
    L := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        L[i] = make([]int, n+1)
    }
    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 || j == 0 {
                L[i][j] = 0
            } else if text1[i-1] == text2[j-1] {
                L[i][j] = L[i-1][j-1] + 1
            } else {
                L[i][j] = max(L[i-1][j], L[i][j-1])
            }
        }
    }
    return L[m][n]
}


func longestCommonSubsequence(text1 string, text2 string) int {
    return longestHelper(text1, text2, len(text1), len(text2))
}
