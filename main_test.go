package main

import (
	"testing"
)

func TestLoopValue(t *testing.T) {
	if LoopValue(6, 0, 10) != 6 {
		t.Error("Loop value 6, min 0, max 10 must be 6!")
	}
	if LoopValue(0, 0, 10) != 0 {
		t.Error("Loop value 0, min 0, max 10 must be 0!")
	}
	if LoopValue(10, 0, 10) != 0 {
		t.Error("Loop value 10, min 0, max 10 must be 0!")
	}
	if LoopValue(11, 0, 10) != 1 {
		t.Error("Loop value 11, min 0, max 10 must be 1!")
	}
	if LoopValue(12, -10, 10) != -8 {
		t.Error("Loop value 12, min -10, max 10 must be -8!")
	}
	if LoopValue(-11, -10, -5) != -6 {
		t.Error("Loop value -11, min -10, max -5 must be -6!")
	}
	if LoopValue(4294967296, 0, 64) != 0 {
		t.Error("Loop value 4294967296, min0, max 64 must be 0!")
	}
}

func TestCompareGenome(t *testing.T) {
	var bot1 Bot
	var bot2 Bot
	var testWorld world

	testWorld.Init()
	bot1.InitBot(0, nil)
	bot2.InitBot(1, nil)

	for i := range bot1.genome {
		bot1.genome[i] = 1
	}
	for i := range bot2.genome {
		bot2.genome[i] = 1
	}

	var result1 = testWorld.compareGenome(&bot1, &bot2)
	if result1 != true {
		t.Error("Identical genome must return true")
	}

	bot2.genome[1] = 2
	var result2 = testWorld.compareGenome(&bot1, &bot2)
	if result2 != true {
		t.Error("One differebnce in genome must return true")
	}

	bot2.genome[2] = 2
	var result3 = testWorld.compareGenome(&bot1, &bot2)
	if result3 == true {
		t.Error("Two or more differences must return false")
	}
}
