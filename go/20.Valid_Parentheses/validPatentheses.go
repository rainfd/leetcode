import "strings"

var in string = "([{"
var out string = ")]}"

func isValid(s string) bool {
    stack := []string{}
    for _, i := range s {
        sign := string(i)
        index := strings.Index(out, sign)
        if index == -1 {
            stack = append(stack, sign)
        } else {
            if len(stack) == 0 {
                return false
            } else if stack[len(stack)-1] != string(in[index]) {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}
