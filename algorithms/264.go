func nthUglyNumber(n int) int {
    ugly := []int{1}
    var i2, i3, i5 int
    
    for len(ugly) < n {
        l := len(ugly)
        for ugly[i2] * 2 <= ugly[l-1] {
            i2++
        }
        for ugly[i3] * 3 <= ugly[l-1] {
            i3++
        }
        for ugly[i5] * 5 <= ugly[l-1] {
            i5++
        }
        ugly = append(ugly, min(ugly[i2] * 2,  ugly[i3] * 3,  ugly[i5] * 5))
    }
    return ugly[n -1]
}

func min(a, b, c int) int {
    if a > b {
        a = b
    }
    if a > c {
        return c
    }
    return a
}
