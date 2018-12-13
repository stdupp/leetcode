func lengthOfLIS(nums []int) int {
    r := 1
    l := len(nums)
    if l == 0 {
        return 0
    }
    dp := make([]int, l+1)
    dp[0] = 1
    for i := 1; i < l; i++ {
        max := 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                tmp := dp[j] + 1
                if tmp > max {
                    max = tmp
                }
            }
        }
        dp[i] = max
        if max > r {
            r = max
        }
    }
    
    return r
}
