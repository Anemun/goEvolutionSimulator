package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var botWorld world

// 1. Init world asd
// 2. Place initial population
// 3. Loop to the end

// For now almost all writes to log are disabled, because of huge impact on performance
// it's not print() function itself who cause slowdown but fmt.Sprint()

func main() {

	// Запуск профайлера ("github.com/pkg/profile")
	// Для вывода результатов надо в терминале запустить go tool pprof .\cpu.pprof
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	rand.Seed(time.Now().UTC().UnixNano()) // this initialize brand new randomizer for current run
	CreateFolders()
	WriteLog("\n---Botworld START---", 1)

	botWorld.Init()
	if debug_fillEntireWorld == true {
		fillEntireWorldWithBots()
	} else {
		placeInitialBots((worldSizeX + worldSizeY) / 5)
	}

	globalTimerStart := time.Now()

	i := 1
	for i < maxTickLimit {
		if botWorld.GetCurrentTickIndex()%100 == 0 {
			elsap := time.Since(globalTimerStart)
			var tps string // ticks per second
			tps = fmt.Sprint(math.Ceil((float64(botWorld.GetCurrentTickIndex())/elsap.Seconds())*100) / 100)

			WriteLog(fmt.Sprint("Processing tick #", botWorld.GetCurrentTickIndex(), " out of ", maxTickLimit, " (", tps, " ticks/s)"), 2)
		}

		botWorld.Tick()
		if debug_checkCollisions == true {
			collisionDetection()
		}
		i++
	}

	FinalSerialization()

	elapsed := time.Since(globalTimerStart)
	fmt.Println("\n---Max bots: ", (worldSizeX * worldSizeY))
	fmt.Println(maxTickLimit, "Ticks took", elapsed)
	fmt.Println(int((float64(maxTickLimit) / elapsed.Seconds())), "ticks per second---")

	WriteLog("\n---Botworld END---", 1)
}

func placeInitialBots(count int) {
	i := 0
	for i < count {
		var newBotCoord = coordinates{rand.Intn(worldSizeX), rand.Intn(worldSizeY)}
		if botWorld.WhatIsOnCoord(newBotCoord, nil) == "empty" {
			i++
			botWorld.NewBot(newBotCoord, nil)
			// WriteLog(fmt.Sprint("Placing starting bot at ", coord), 4)
		}
	}
}

func fillEntireWorldWithBots() {
	i, j := 0, 0
	for i < worldSizeX {
		for j < worldSizeY {
			var newBotCoord = coordinates{i, j}
			botWorld.NewBot(newBotCoord, nil)
			// WriteLog(fmt.Sprint("Placing starting bot at ", coord), 4)
			j++
		}
		j = 0
		i++
	}
}

func placeTestBots() {
	botWorld.NewBot(coordinates{3, 3}, nil)
	botWorld.bots[3][3].genome[0] = 10
	botWorld.bots[3][3].genome[1] = 0
	botWorld.bots[3][3].genome[3] = 61

	botWorld.NewBot(coordinates{7, 7}, nil)
	botWorld.bots[7][7].genome[0] = 10
	botWorld.bots[7][7].genome[1] = 2
	botWorld.bots[7][7].genome[3] = 61
}

func collisionDetection() {
	var i, j int = 0, 0
	for i < worldSizeX {
		for j < worldSizeY {
			if botWorld.bots[i][j] != nil && botWorld.food[i][j] != nil {
				panic("There is two object on same tile!")
			}
			j++
		}
		i++
	}
}
