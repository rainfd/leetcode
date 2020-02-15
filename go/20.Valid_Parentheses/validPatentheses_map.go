import "strings"

var in string = "([{"
var out string = ")]}"

func isValid(s string) bool {
    dict := map[string]string{
        ")": "(",
        "]": "[",
        "}": "{",
    }
    top := "N"
    stack := []string{}
    for _, i := range s {
        sign := string(i)

        if len(stack) == 0 {
            top = "N"
        } else {
            top = stack[len(stack)-1]
        }

        value, ok := dict[sign]
        // ({[
        if !ok {
            stack = append(stack, sign)
            continue
        }
        // )}]
        if top != value {
            return false
        } else {
            // )}]
            stack = stack[:len(stack)-1]
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}
