func intToRoman(num int) string {
    var s string
    m := num / 1000
    s = construct(s, m, "M")
    num = num % 1000
    if num >= 900 {
        s = construct(s, 1, "CM")
        num = num - 900
    } else if num >= 500 {
        s = construct(s, 1, "D")
        num = num - 500
    } else if num >= 400 {
        s = construct(s, 1, "CD")
        num = num - 400
    }
    
    if num >= 100 {
        c := num / 100
        s = construct(s, c, "C")
        num = num % 100
    } 
    
    if num >= 90 {
        s = construct(s, 1, "XC")
        num = num - 90
    } else if num >= 50 {
        s = construct(s, 1, "L")
        num = num - 50
    } else if num >= 40 {
        s = construct(s, 1, "XL")
        num = num - 40
    }
    if num >= 10 {
        x := num / 10
        s = construct(s, x, "X")
        num = num % 10
    }
    
    if num == 9 {
        s = construct(s, 1, "IX")
        return s
    } else if num >= 5 {
        s = construct(s, 1, "V")
        num = num - 5
    } else if num >= 4 {
        s = construct(s, 1, "IV")
        num = num -4
    }
    
    s = construct(s, num, "I")
    
    return s
}

func construct(s string, n int, e string) string {
    var tmp string
    for i := 0; i < n; i++ {
        tmp += e
    }
    return s + tmp
}
