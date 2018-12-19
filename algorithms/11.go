func maxArea(height []int) int {
    var max int
    
    i := 0
    j := len(height) - 1
    
    for i < j {
        lower := height[i]
        if height[j] < lower {
            lower = height[j]
        }
        area := (j - i) * lower
        if area > max {
            max = area
        }
        
        if height[i] > height[j] {
            j--
        } else {
            i++
        }
    }
    
    return max
}
