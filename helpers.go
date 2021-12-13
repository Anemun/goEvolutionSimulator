package main

// LoopValue returns value, looped between min (included) and max (not included)
// LoopValue for example, val=6, min=0, max=10 returns 6
// LoopValue for example, val=0, min=0, max=10 returns 0
// LoopValue for example, val=10, min=0, max=10 returns 0
// LoopValue for example, val=11, min=0, max=10 returns 1
// LoopValue for example, val=12, min=-10, max=10 returns -8
// LoopValue for example, val=-11, min=-10, max=-5 returns -6
func LoopValue(value int, min int, max int) int {
	//fmt.Println(fmt.Sprint("Looping: value=", value, ", min=", min, ", max =", max))
	for value < min {
		diff := min - value
		value = max - diff
	}

	for value >= max {
		diff := value - max
		value = min + diff
	}

	//fmt.Println(fmt.Sprint("Result: ", value))
	return value
}
