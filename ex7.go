import "math"

func rn(p, n int) int {
    if n <= 2 {
        return (p + 99*(p % 10)) / 10;
    }
    exp := int(math.Pow10(n-1))
    return 10*(rn(p % exp, n-1)) + (p - (p % exp)) / exp;
}


func reverse(x int) int {
    max := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
    cpy := x
    n := 0
    for cpy != 0 {
        cpy /= 10
        n += 1
    }
    if n >= 10 {
        if n > 10 {
            return 0
        }
        cpy = x
        r := 0
        overflow := false
        for cpy != 0 && r < len(max) {
            digit := cpy % 10
            if digit == max[r] || digit == -1*max[r] {
                cpy /= 10
                r += 1
                continue
            } else if digit > max[r] || (digit < -1*max[r] || digit < -8) {
                overflow = true
                break
            } else if digit < max[r] || (digit > -1*max[r] || digit > -8) {
                break
            }
            r += 1
            cpy /= 10
        }
        if overflow {
            return 0
        }
    } else if n == 1 {
        return x
    }
    return rn(x, n)
}
