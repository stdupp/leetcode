func divide(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    if dividend == -2147483648 && divisor == -1 {
        return 2147483647
    }

    var sign int
    if dividend < 0 {
        sign -= 1
        dividend = -dividend
    } else {
        sign += 1
    }
    
    if divisor < 0 {
        sign -= 1
        divisor = -divisor
    } else {
        sign += 1
    }
    
    var r int
    for dividend >= divisor {
        r++
        dividend -= divisor
    }
    
    if sign == 0 {
        r = -r
    }
    return r
}
