func isValid(s string) bool {
    var stack []int
    
    for _, tmp := range s {
        v := int(tmp)
        if v == int('{') || v == int('(') || v == int('[') {
            stack = append(stack, v)
            continue
        }
        
        if len(stack) == 0 {
            return false
        }
    
        last := stack[len(stack) -1]
        if last - v !=  {
            return false
        } else {
            stack = stack[:len(stack) - 1]
        }
        continue
        
    
        if v == '}' {
            last := stack[len(stack) -1]
            if last != '{' {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
            continue
        }
        
        if v == ']' {
            last := stack[len(stack) -1]
            if last != '[' {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
            continue
        }
        
        if v == ')' {
            last := stack[len(stack) -1]
            if last != '(' {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
            continue
        }
    }
    
    if len(stack) != 0 {
        return false
    }
    
    return true
}
