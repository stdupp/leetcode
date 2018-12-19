func numTrees(n int) int {
    t := make([]int, n+1)
    t[0], t[1] = 1, 1
    if n <= 1 {
        return n
    }
    
    for i:= 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            t[i] += t[j-1] * t[i-j]
        }
    }
    
    return t[n]
}
