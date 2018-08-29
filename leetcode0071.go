func simplifyPath(path string) string {
    stack := make([]string, 0)
    for _, s := range strings.Split(path, "/") {
        if s == "" || s == "." {
            continue
        }
        if s == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }
        } else {
            stack = append(stack, s)
        }
    }
    return "/" + strings.Join(stack, "/")
}