func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    var flagr, flagc bool
    
    r := len(obstacleGrid)
    c := len(obstacleGrid[0])
    
    path := make([][]int, r)
    for i := range path {
        row := make([]int, c)
        path[i] = row
    }
    
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if i == 0 {
                if flagr == true {
                    path[i][j] = 0
                } else if obstacleGrid[i][j] == 1 {
                    path[i][j] = 0
                    flagr = true
                } else {
                    path[i][j] = 1
                }
            } else if j == 0 {
                if flagc == true {
                    path[i][j] = 0
                } else if obstacleGrid[i][j] == 1 {
                    path[i][j] = 0
                    flagc = true
                } else {
                    path[i][j] = 1
                }
            } else if obstacleGrid[i][j] == 1 {
                path[i][j] = 0
            } else {
                path[i][j] = path[i -1][j] + path[i][j-1]   
            }
        }
    }
    
    return path[r -1][c -1]
}
