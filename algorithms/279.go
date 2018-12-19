func numSquares(n int) int {
    dp := make([]int, n+1)
    for i:=1; i <= n; i++ {
        count := 0
        for j := 1;j*j <= i; j++ {
            tmp := dp[i-j*j] + 1
            if tmp < count{
                count = tmp
            } else if count == 0 {
                count = tmp
            }
        }
        dp[i] = count
    }
    
    return dp[n]
}
