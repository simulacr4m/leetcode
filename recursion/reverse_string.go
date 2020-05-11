func helper(index int, s []byte) {
    if index == len(s) / 2 {
        return
    }
    tmp := s[index]
    s[index] = s[len(s)-index-1]
    s[len(s)-index-1] = tmp
    helper(index+1, s)
}

func reverseString(s []byte) {
    helper(0, s)
}
