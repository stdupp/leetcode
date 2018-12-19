func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n; i++ {
        doRotate(matrix, i, n-1-i)
    }
}

func doRotate(matrix [][]int, start, end int) {
    for i:= start; i < end; i++ {
        tmp := matrix[i][end]
        matrix[i][end] = matrix[start][i]
        matrix[start][i] = matrix[end+start-i][start]
        matrix[end+start-i][start] = matrix[end][end+start-i]
        matrix[end][end+start-i] = tmp
    }
}
