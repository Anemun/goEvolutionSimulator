package main

import (
	"math/rand"
)

// Bot class
type Bot struct {
	index                  uint64
	coordX                 int
	coordY                 int
	energy                 int
	genome                 []byte
	organs                 []*Organ
	commandPointer         int
	isDead                 bool
	doNextMinorCommand     bool
	minorCommandCount      int
	majorCommandPointsLeft int
	age                    int
	carnivoreRating        int
	herbivoreRating        int
}

func (bot *Bot) AddCarnivoreRating(increment int) {
	bot.carnivoreRating += increment
	if bot.carnivoreRating > 100 {
		bot.carnivoreRating = 100
	}
	if bot.carnivoreRating < 0 {
		bot.carnivoreRating = 0
	}
}
func (bot *Bot) AddHerbivoreRating(increment int) {
	bot.herbivoreRating += increment
	if bot.herbivoreRating > 100 {
		bot.herbivoreRating = 100
	}
	if bot.herbivoreRating < 0 {
		bot.herbivoreRating = 0
	}
}

// SetCommandPointer SET
func (bot *Bot) SetCommandPointer(newPointer int) {
	bot.commandPointer = LoopValue(newPointer, 0, botGenomeSize)
}

// IncrementCommandPointer ++
func (bot *Bot) IncrementCommandPointer(increment int) {
	bot.SetCommandPointer(bot.CommandPointer() + increment)
}

// CommandPointer GET
func (bot *Bot) CommandPointer() int {
	return bot.commandPointer
}

// CalculateMaxEnergy - return maximum bot energy based on organs count
func (bot *Bot) CalculateMaxEnergy() int {
	return maxBaseBotEnergy + (len(bot.organs) * maxBaseBotEnergyPerOrgan)
}

// AddEnergy SET
func (bot *Bot) AddEnergy(increment int) {
	bot.energy += increment
	botMaxEnergy := bot.CalculateMaxEnergy()
	if bot.energy > botMaxEnergy {
		if makeChildIfEnergySurplus == true {
			bot.commandCHILD()
		} else {
			bot.energy = botMaxEnergy
		}
	}
	if bot.energy < 0 {
		bot.energy = 0
	}
}

// InitBot initializes new bot
func (bot *Bot) InitBot(index uint64, parent *Bot) {
	bot.SetCommandPointer(0)
	bot.isDead = false
	bot.index = index
	bot.genome = make([]byte, botGenomeSize)
	if parent == nil {
		bot.energy = botGenomeSize
		for i := range bot.genome {
			bot.genome[i] = byte(rand.Intn(botGenomeSize))
		}
	}

	if parent != nil {
		copy(bot.genome, parent.genome)
		parent.genome[0] = 99
		var mutate = rand.Float64()
		if mutate <= mutateChance {
			var mutateByte int
			mutateByte = LoopValue(rand.Intn(botGenomeSize), 0, botGenomeSize)
			bot.genome[mutateByte] = byte(LoopValue(rand.Intn(botGenomeSize), 0, botGenomeSize))
		}
		bot.energy = int(float64(parent.energy) * childEnergyFraction)
	}
}

func (bot *Bot) doCommand() {
	//Bot commands defined in botCommands.go
	switch bot.genome[bot.commandPointer] {
	case 0:
		bot.commandSTAY()
	case 5:
		bot.commandLOOKa()
	case 10:
		bot.commandMOVEa()
	case 15:
		bot.commandEATa()
	case 20:
		bot.commandPHOTOSYNTESIS()
	// case 25:
	// 	bot.commandORGAN()
	case 30:
		bot.commandCHILD()
	default:
		bot.forwardPointer()
	}
}

// Tick bot logic
func (bot *Bot) Tick() {
	if botWorld.bots[bot.coordX][bot.coordY] == nil {
		panic("why bot nil!?")
	}

	if bot.energy <= 0 {
		botWorld.BotIsDead(bot)
		return
	}

	if bot.age >= oldAgeDyingCap {
		botWorld.BotIsDead(bot)
		return
	} else {
		bot.age++
	}

	if bot.isDead {
		return
	}

	if botWorld.GetCurrentTickIndex()%10 == 0 {
		bot.AddCarnivoreRating(-1)
		bot.AddHerbivoreRating(-1)
	}

	bot.AddEnergy(-1 * (botEnergyTickCost + len(bot.organs)*botEnergyTickCostPerOrgan))
	bot.minorCommandCount = 0
	bot.doNextMinorCommand = true
	bot.majorCommandPointsLeft = initialMajorCommandPointsPerTurn

	for i := range bot.organs {
		bot.organs[i].tick()
	}

	bot.doCommand()

	// // Бот имеет несколько больших действий, в каждом есть энное количество маленьких
	// for bot.majorCommandPointsLeft >= majorCommandPointsCostPerAcrion { // пока остались очки на хотя бы одно большое действие
	// 	for bot.doNextMinorCommand == true &&
	// 		bot.minorCommandCount < maxMinorCommandsPerMajorCommand { // делать маленькие действия
	// 		bot.minorCommandCount++
	// 		bot.doCommand()
	// 	}
	// 	bot.majorCommandPointsLeft -= majorCommandPointsCostPerAcrion
	// 	bot.minorCommandCount = 0
	// 	bot.doNextMinorCommand = true
	// }

	// botWorld.setBotOnCoord(Coord{bot.CoordX + 1, bot.CoordY + 1}, *bot)
	// WriteLog(fmt.Sprint("Tick ", botWorld.GetCurrentTickIndex(), ": bot ", bot.index, " ", "is on tile ", bot.coordX, ".", bot.coordY, ". Bot energy: ", bot.energy, ". Bot pointer: [", bot.commandPointer, "]: ", bot.genome[bot.commandPointer]), 4)
}

// Take value from next genome byte and make a direction out of it (default direction count is 8 so value from 0 to 7, where 0 is up, 1 is up-right and 7 is up-left)
func (bot *Bot) getAbsoluteDirection() int {
	var direction int
	cp := LoopValue(bot.CommandPointer()+1, 0, botGenomeSize)
	direction = int(bot.genome[cp] % directionsCount)
	return direction
}

func (bot *Bot) getAdjascentCoordByDirection(direction int) coordinates {
	if directionsCount != 8 {
		panic("This function covers only 8 directions! Update it if you change directions count")
	}

	var targetCoord = coordinates{bot.coordX, bot.coordY}
	switch direction {
	case 0:
		targetCoord.y++
	case 1:
		targetCoord.x++
		targetCoord.y++
	case 2:
		targetCoord.x++
	case 3:
		targetCoord.x++
		targetCoord.y--
	case 4:
		targetCoord.y--
	case 5:
		targetCoord.x--
		targetCoord.y--
	case 6:
		targetCoord.x--
	case 7:
		targetCoord.x--
		targetCoord.y++
	}
	targetCoord.x = LoopValue(targetCoord.x, 0, worldSizeX)
	targetCoord.y = LoopValue(targetCoord.y, 0, worldSizeY)

	return targetCoord
}
