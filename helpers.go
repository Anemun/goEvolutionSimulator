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

package main

import (
	"fmt"
)

func main() {
	LoopValue(0, 0, 10)	// must be 0
	LoopValue(-1, 0, 10)	// must be 9
	LoopValue(2, 0, 10)	// must be 2
	LoopValue(11, 0, 10)	// must be 1
	LoopValue(10, 0, 10)	// must be 0
	LoopValue(4, 5, 10)	// must be 14
	LoopValue(-5, -5, 5)	// must be -5
	LoopValue(0, -5, 5)	// must be -5
	LoopValue(-2, -5, 5)	// must be -2
	LoopValue(-6, -5, 5)	// must be -1
}

func LoopValue(value int, min int, capacity int) int {
	fmt.Println(fmt.Sprint("Looping: value=", value, ", min=", min, ", capacity =", capacity ))
	while value < min {
    diff := min - value
    value = min+capacity-diff
  }

  while value >= min+capacity {
    diff := value-max
    value = min + diff
  }


	fmt.Println(fmt.Sprint("Result: ", value))
	return value
}
