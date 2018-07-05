package main

// LoopValue returns value, looped between min (included) and max (not included)
// LoopValue for example, val=6, min=0, max=10 returns 6
// LoopValue for example, val=0, min=0, max=10 returns 0
// LoopValue for example, val=10, min=0, max=10 returns 0
// LoopValue for example, val=11, min=0, max=10 returns 1
// LoopValue for example, val=12, min=-10, max=10 returns -8
func LoopValue(value int, min int, max int) int {
	// WriteLog(fmt.Sprint("Looping: value=", value, ", min=", min, ", max=", max), 5)
	if min == max {
		return min
	}

	if value < min {
		value = max - (value % max)
		// value = max + (value - min)
	}

	if value >= max {
		value = min + (value % max)
		// value = value - (max - min)
	}

	//WriteLog(fmt.Sprint("Result: ", value), 5)
	return value
}
