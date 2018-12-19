func myAtoi(str string) int {
    var sign int
    var value int
    for _, v := range str {
        if v == '-' {
            if sign != 0 {
                break
            }
            sign = -1
            continue
        }
        
        if v == '+' {
            if sign != 0 {
                break
            }
            sign = 1
            continue
        }
        
        if v >= '0' && v <= '9' {
            if sign == 0 {
                sign = 1
            }
            value = value * 10 + int(v -'0')
            if value > 2147483648 {
                break
            }
            continue
        }
        
        if v == ' ' {
            if value != 0 || sign !=0 {
                break
            }
            continue
        }
        
        break
    }
    

    
    value = value * sign
    
    if value <  -2147483648 {
        return -2147483648
    }
    if value > 2147483647 {
        return 2147483647
    }
    
    return value
    
}
