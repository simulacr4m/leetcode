import "math"

func isPerfectSquare(num int) bool {
    n := float64(num)
    x := n
    for i := 0; i < 1000; i++ {
        x = 0.5*(x + n/x)
    }
    return math.Floor(x) == math.Ceil(x)
}
