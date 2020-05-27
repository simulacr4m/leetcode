func addDigits(num int) int {
    if num < 10 {
        return num
    } else if num % 9 == 0 {
        return 9
    }
    return num - 9*(num / 9)
}
