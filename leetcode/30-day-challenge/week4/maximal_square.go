func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    R, C := len(matrix), len(matrix[0])
    S := make([][]int, R)
    for i := 0; i < R; i++ {
        S[i] = make([]int, C)
    }
    for i := 0; i < R; i++ {
        if matrix[i][0] == '1' {
            S[i][0] = 1
        } else {
            S[i][0] = 0
        }
    }
    
    for i := 0; i < C; i++ {
        if matrix[0][i] == '1' {
            S[0][i] = 1
        } else {
            S[0][i] = 0
        }
    }
    
    for i := 1; i < R; i++ {
        for j := 1; j < C; j++ {
            if matrix[i][j] == '1' {
                S[i][j] = min(S[i][j-1], min(S[i-1][j], S[i-1][j-1])) + 1
            } else {
                S[i][j] = 0
            }
        }
    }
    
    maxS := S[0][0]
    
    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            if maxS < S[i][j] {
                maxS = S[i][j]
            }
        }
    }
    return maxS*maxS
}
