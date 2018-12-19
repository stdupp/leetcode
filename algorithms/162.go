func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    start := 0
    end := len(nums) - 1
    for start <= end {
        mid := (start+end)/2
        
        if mid >0 && nums[mid] < nums[mid-1] {
            end = mid -1
        } else if mid < len(nums)-1 && nums[mid] < nums[mid+1]{
            start = mid + 1
        } else {
            return mid
        }
    }
    
    return start
}
