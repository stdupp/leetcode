/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "sort"

func merge(intervals []Interval) []Interval {
    if len(intervals) <= 1 {
        return intervals
    }
    sort.Slice(intervals, func(i, j int) bool {return intervals[i].Start < intervals[j].Start})
    
    var result []Interval
    current := intervals[0]
    for i := 1; i < len(intervals); i++ {
        if current.End >= intervals[i].Start {
            if current.End < intervals[i].End {
                current.End = intervals[i].End
            }
        } else {
            result = append(result, current)
            current = intervals[i]
        }
    }
    
    result = append(result, current)
    
    return result
}
