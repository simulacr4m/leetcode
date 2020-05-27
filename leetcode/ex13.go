func romanToInt(s string) int {
    numerals := map[string]int {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    value := 0
    for i := 0; i < len(s); i++ {
        j := numerals[s[i:i+1]]
        k := 0
        if i < len(s)-1 {
            k = numerals[s[i+1:i+2]]
        }
        if j < k {
            value += k - j
            i += 1
        } else {
            value += j
        }
    }
    return value
}
