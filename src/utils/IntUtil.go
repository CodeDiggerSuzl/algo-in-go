package utils

// get the max int value
func MaxInt(a, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}

func MinInt(a, b int) int {
    if a >= b {
        return b
    } else {
        return a
    }
}
