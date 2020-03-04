import "math"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    palindrome, n, i := 0, 0, 1
    for cpy := x; cpy != 0; cpy /= 10 {
        n += 1
    }
    i = n - 1
    for cpy := x; cpy != 0; cpy /= 10 {
        digit := cpy % 10
        palindrome += digit * int(math.Pow10(i))
        i -= 1
    }
    return palindrome == x
}
