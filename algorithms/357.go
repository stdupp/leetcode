func countNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    dp := make([]int, n)
    dp[0] = 10
    for i := 1; i < n;i++ {        
        dp[i] = dp[i-1] + f(i+1)
    }
    
    return dp[n-1]
    
}

func f(n int) int {
    if n >= 11 {
        return 0
    }
    total := 9
    for i :=0; i < n-1; i++ {
        total *= 9 - i
    }
    
    return total
}
