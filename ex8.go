func isValid(str string) int {
    if len(str) == 0 {
        return 0
    }
    digitStart := 0
    if str[0] == '+' || str[0] == '-' {
        digitStart += 1
    }
    for digitStart < len(str) && unicode.IsDigit(rune(str[digitStart])) {
        digitStart += 1
    }
    if (digitStart > 11) {
        if str[0] == '-' {
            return math.MinInt32
        } else {
            return math.MaxInt32
        }
    }
    return -1
}

func makeValidStr(str string) string {
    str = strings.TrimSpace(str)
    if len(str) == 0 {
        return str
    }
    x, i := 0, 0
    sign := ""
    if str[0] == '-' || str[0] == '+' {
        i += 1
        x += 1
        sign = string(str[0])
    }
    for ; i < len(str) && unicode.IsDigit(rune(str[i])); {
        i += 1
    }
    str = str[0:i]
    for ; x < len(str) && str[x] == '0'; {
        x += 1
    }
    str = str[x:]
    str = sign + str
    return str
}


func myAtoi(str string) int {
    var res int64
    var sign int64
    str = makeValidStr(str)
    x := isValid(str)
    if x != -1 {
        return x
    }
    if str[0] == '-' {
        sign = -1
    } else {
        sign = 1
    }
    var r rune
    for i := 0; i < len(str); i++ {
        r = rune(str[i])
        if r == '+' || r == '-' {
            continue
        }
        if unicode.IsDigit(r) {
            res = 10*res + (int64(r)-'0')
        } else {
            break
        }
    }
    res *= sign
    if res > math.MaxInt32 {
        return math.MaxInt32
    }
    if res < math.MinInt32 {
        return math.MinInt32
    }
    return int(res)
}
