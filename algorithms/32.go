func longestValidParentheses(s string) int {
    num := 0
    max := 0
    length :=0
    index := 0
    for i:=0; i < len(s);{
        b := byte(s[i])
        if b == '(' {
            num++
            length++
            i++
        } else {
            num--
            if num < 0 {
                num =0
                length = 0
                index++
                i = index
            } else {
                length++
                if num == 0 {
                    if length > max {
                        max = length
                    }
                }
                i++
            }
        }
        if i == len(s) && num !=0 {
            index++
            i=index
            num =0
            length = 0
        }
        
    }
    return max
    
}
