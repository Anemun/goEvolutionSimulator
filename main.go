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

	state := "Running"

	rand.Seed(time.Now().UTC().UnixNano()) // this initialize brand new randomizer for current run
	if serializationEnabled == true {
		createFolders()
	}
	WriteLog("\n---Botworld START---", 1)

	botWorld.Init()
	if debugfillEntireWorld == true {
		fillEntireWorldWithBots()
	} else if debugPlaceCustomBots == true {
		placeTestBots()
	} else {
		placeInitialBots(initialBotCount)
	}

	globalTimerStart := time.Now()

	i := 1
	for i < maxTickLimit && state == "Running" {
		if botWorld.GetCurrentTickIndex()%100 == 0 {
			elsap := time.Since(globalTimerStart)
			var tps string // ticks per second
			tps = fmt.Sprint(math.Ceil((float64(botWorld.GetCurrentTickIndex())/elsap.Seconds())*100) / 100)

			WriteLog(fmt.Sprint("Processing tick #", botWorld.GetCurrentTickIndex(), " out of ", maxTickLimit, " (", tps, " ticks/s)", " Bots alive: ", aliveBotCount), 2)
		}

		botWorld.Tick()
		if debugcheckCollisions == true {
			collisionDetection()
		}
		i++
	}

	if serializationEnabled == true {
		FinalSerialization()
	}

	elapsed := time.Since(globalTimerStart)
	fmt.Println("\n---Max possible bots: ", (worldSizeX * worldSizeY))
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
			WriteLog(fmt.Sprint("Placing starting bot at ", newBotCoord), 4)
		}
	}
}

// TEST CODE
func fillEntireWorldWithBots() {
	i, j := 0, 0
	for i < worldSizeX {
		for j < worldSizeY {
			var newBotCoord = coordinates{i, j}
			botWorld.NewBot(newBotCoord, nil)
			WriteLog(fmt.Sprint("Placing starting bot at ", newBotCoord), 4)
			j++
		}
		j = 0
		i++
	}
}

// TEST CODE
func placeTestBots() {
	botWorld.NewBot(coordinates{5, 5}, nil)
	botWorld.bots[5][5].energy = 63
	botWorld.bots[5][5].genome[0] = 25
	botWorld.bots[5][5].genome[10] = 10
	botWorld.bots[5][5].genome[11] = 0
	botWorld.bots[5][5].genome[12] = 62

	botWorld.NewBot(coordinates{5, 8}, nil)
	botWorld.bots[5][8].energy = 63
	botWorld.bots[5][8].genome[0] = 0
	botWorld.bots[5][8].genome[1] = 0
	botWorld.bots[5][8].genome[3] = 0
}

func collisionDetection() {
	var i, j int = 0, 0
	for i < worldSizeX {
		for j < worldSizeY {
			if botWorld.bots[i][j] != nil {
				if botWorld.food[i][j] != nil {
					panic("There is two object on same tile!")
				}
				if botWorld.organs[i][j] != nil {
					panic("There is two object on same tile!")
				}
			}
			if botWorld.food[i][j] != nil {
				if botWorld.organs[i][j] != nil {
					panic("There is two object on same tile!")
				}
			}
			j++
		}
		i++
	}
}
