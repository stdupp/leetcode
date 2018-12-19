func wordBreak(s string, wordDict []string) bool {
    slen := len(s)
    dp := make([]bool, slen+1)
    dp[0] = true
    for i := 1; i <= slen; i++ {
        for _, word := range wordDict {
            size := len(word)
            if i < size {
                continue
            }
            
            if dp[i-size] && s[i-size:i] == word {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[slen]
}

