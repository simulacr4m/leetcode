func isValid(s string) bool {
    if s == "" {
        return true
    }
    var stack []string
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, string(s[i]))
        } else {
            if len(stack) <= 0 {
                return false
            }
            char := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if (char == "(" && s[i] != ')') || (char == "{" && s[i] != '}') || (char == "[" && s[i] != ']') {
                return false
            }
        }
    }
    return len(stack) == 0
}
