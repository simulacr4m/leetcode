func isOneBitCharacter(bits []int) bool {
    for i := 0; i < len(bits)-1; {
        if bits[i] == 1 {
            if i+1 == len(bits)-1 {
                return false
            }
            i += 2
        } else {
            i += 1
        }
    }
    return true
}
