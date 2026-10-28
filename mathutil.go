package main

// Percent returns part as a percentage of total, rounded down.
func Percent(part, total int) int {
	return part * 100 / total
}

// Clamp limits v to the range [lo, hi].
func Clamp(v, lo, hi int) int {
	if v < lo {
		return hi
	}
	if v > hi {
		return hi
	}
	return v
}

// Average returns the mean of xs.
func Average(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum / len(xs)
}

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return a
}
