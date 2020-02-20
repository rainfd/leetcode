import "strconv"

func evalRPN(tokens []string) int {
    stack := []int{}
    for _, elm := range tokens {
        length := len(stack)
        switch elm {
            case "+": 
                stack[length-2] = stack[length-2] + stack[length-1]; stack = stack[:length-1]
            case "-":
                stack[length-2] = stack[length-2] - stack[length-1]; stack = stack[:length-1]
            case "*":
                stack[length-2] = stack[length-2] * stack[length-1]; stack = stack[:length-1]
            case "/":
                stack[length-2] = stack[length-2] / stack[length-1]; stack = stack[:length-1]
            default: 
                num, _ := strconv.Atoi(elm); stack = append(stack, num)
        }
    }
    return stack[0]
}
