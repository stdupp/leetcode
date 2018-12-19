func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    start := 0
    end := len(nums) -1
    
    for start <= end {
        mid := (start + end)/2
        if nums[mid] == target {
            return mid
        }
        
        if nums[mid] >= nums[start] {
            if nums[mid] >= target && target >= nums[start] {
                end = mid -1
            } else {
                start = mid + 1
            }
        } else {
            if nums[mid] <= target && target <= nums[end] {
                start = mid + 1
            } else {
                end = mid -1
            }
        }
    }
    
    return -1
}
