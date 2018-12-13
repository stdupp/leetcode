func integerBreak(n int) int {
    dp := make([]int, n+1)
    dp[1] = 1

    for i := 2; i <= n; i++ {
        var max int
        for j := 1; j < i;j++ {
            a1 := j
            a2 := i-j
            if j < dp[j] {
                a1 = dp[j]
            }
            if i-j < dp[i-j] {
                a2 = dp[i-j]
            }
            tmp := a1 * a2
            if tmp > max {
                max = tmp
            }
        }
        dp[i] = max
    }
    
    return dp[n]
}
