func uniquePaths(m int, n int) int {
    if m > n {
        return uniquePaths(n, m)
    }
    
    cur := make([]int, m)
    for i := range cur {
        cur[i] = 1
    }
    
    for i:= 1; i < n; i++ {
        for j := 1;j < m;j++ {
            cur[j] += cur[j -1]
        }
    }
    
    return cur[m - 1]
}
