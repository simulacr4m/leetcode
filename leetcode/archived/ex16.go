func addBinary(a string, b string) string {
    alen, blen := len(a), len(b)
    if alen < blen {
        for i := 0; i < blen-alen; i++ {
            a = "0" + a
        }
    } else {
        for i := 0; i < alen-blen; i++ {
            b = "0" + b
        }
    }
    leading := 0
    carry := false
    c := ""
    for i := len(a)-1; i >= 0; i-- {
        if (a[i] == '0' && b[i] == '1') || (a[i] == '1' && b[i] == '0') {
            if carry {
                c = "0" + c
            } else {
                c = "1" + c
            }
        } else if a[i] == '0' && b[i] == '0' {
            if carry {
                c = "1" + c
                carry = false
            } else {
                c = "0" + c
            }
        } else if a[i] == '1' && b[i] == '1' {
            if carry {
                c = "1" + c
            } else {
                c = "0" + c
                carry = true
            }
        }
    }
    if carry {
        c = "1" + c
    }
    if len(c) > 1 {
        for ; leading < len(c) && c[leading] == '0'; {
            leading += 1
        }
    }
    return c[leading:]
}
