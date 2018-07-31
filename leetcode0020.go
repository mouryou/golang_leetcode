func isValid(s string) bool {
    symbols := map[rune]byte {')':'(', ']':'[', '}':'{'}
    stack := list.New()
    for _, c := range s {
        if _, ok := symbols[c]; ok {
            if stack.Len() == 0 {
                return false
            }
            if stack.Back().Value.(int32) == int32(symbols[c]) {
                stack.Remove(stack.Back())
            } else {
                return false
            }
        } else {
            stack.PushBack(c)
        }
    }
    return stack.Len() == 0
}