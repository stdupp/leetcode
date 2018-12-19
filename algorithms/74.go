func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    
    bottomRow := 0
    topRow := len(matrix) -1
    c := len(matrix[0]) -1
    
    for bottomRow <= topRow {
        row := (bottomRow + topRow) /2
        if matrix[row][0] <= target && matrix[row][c] >= target {
            return binarySearch(matrix[row], target)
        } else if matrix[row][0] > target {
            topRow = row -1
            
        } else {
            bottomRow = row + 1
        }
    }
    
    return false
}

func binarySearch(row []int, target int ) bool {
    start := 0
    end := len(row) -1
    for start <= end {
        i := (start + end) /2
        if row[i] == target {
            return true
        }else if row[i] < target {
            start = i + 1
        } else {
            end = i -1 
        }
    }
    
    return false
}
