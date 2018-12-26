func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }
    
    if len(s) == 1 {
        return 1
    }
    
    dp := []int{1, 1}
    for j:=1; j < len(s); j++ {
        a := string(s[j-1])
        b := string(s[j])
        res := 0
        if !isValid(b) && !isValid(a+b) {
            return 0
        }
        
        if isValid(b) {
            res += dp[1]
        }
        
        if isValid(a+b) {
            res += dp[0]
        }
        
        dp[0], dp[1] = dp[1], res
    }
    return dp[1]
}

func isValid(s string) bool {
    if s[0] == '0' {
        return false
    }
    
    i, _ := strconv.Atoi(s)
    return i >= 1 && i <= 26
}
