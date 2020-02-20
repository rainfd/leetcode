import "strings"

func simplifyPath(path string) string {
    dirs := strings.Split(path, "/")
    stack := []string{}
    for _, dir := range dirs {
        switch dir {
            case "", ".":
                continue
            case "..":
                if len(stack) > 0 {
                    stack = stack[:len(stack)-1]
                }
            default:
                stack = append(stack, dir)
        }
    }
    
    return "/" + strings.Join(stack, "/")
}
