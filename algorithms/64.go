func minPathSum(grid [][]int) int {
    row := len(grid)
    col := len(grid[0])
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if i == 0 {
                if j > 0 {
                    grid[i][j] += grid[i][j-1]
                }
            } else if j == 0 {
                if i > 0 {
                    grid[i][j] += grid[i-1][j]
                }
            } else {
                if grid[i-1][j] < grid[i][j-1] {
                    grid[i][j] += grid[i-1][j]
                } else {
                    grid[i][j] += grid[i][j-1]
                }
            }
        }
    }
    
    return grid[row-1][col-1]
}
