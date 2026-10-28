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
