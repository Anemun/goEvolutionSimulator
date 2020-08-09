package main

import (
	"math/rand"
)

// Bot class
type Bot struct {
	index              uint64
	coordX             int
	coordY             int
	energy             int
	genome             []byte
	organs             []*Organ
	commandPointer     int
	isDead             bool
	doNextMinorCommand bool
	minorCommandCount  int
	majorCommandPointsLeft   int
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
	if parent == nil {
		bot.energy = botGenomeSize
		bot.genome = make([]byte, botGenomeSize)
		for i := range bot.genome {
			bot.genome[i] = byte(rand.Intn(botGenomeSize))
		}
	}

	if parent != nil {
		bot.energy = 0
		bot.AddEnergy(int(float64(parent.energy) * childEnergyFraction))
		bot.genome = parent.genome
		var mutate = rand.Float64()
		if mutate <= mutateChance {
			var mutateByte int
			mutateByte = LoopValue(rand.Intn(botGenomeSize), 0, botGenomeSize)
			bot.genome[mutateByte] = byte(LoopValue(rand.Intn(botGenomeSize), 0, botGenomeSize))
		}
	}
}

func (bot *Bot) doCommand() {
	//Bot commands defined in botCommands.go
	switch bot.genome[bot.commandPointer] {
	case 0:
		bot.commandSTAY()
	case 5:
		bot.commandLOOKa()
	// case 10:
	// 	bot.commandMOVEa()
	// case 15:
	// 	bot.commandEAT()
	case 20:
		bot.commandPHOTOSYNTESIS()
	case 25:
		bot.commandORGAN()
	// case 30:
	// 	bot.commandCHILD()
	default:
		bot.forwardPointer()
	}
}

// Tick bot logic
func (bot *Bot) Tick() {
	if bot.isDead {
		return
	}

	bot.AddEnergy(-1 * (botEnergyTickCost + len(bot.organs)*botEnergyTickCostPerOrgan))
	bot.minorCommandCount = 0
	bot.doNextMinorCommand = true
	bot.majorCommandPointsLeft = initialMajorCommandPointsPerTurn

	for i := range bot.organs {
		bot.organs[i].tick()
	}

	// Бот имеет несколько больших действий, в каждом есть энное количество маленьких
	for bot.majorCommandPointsLeft >= majorCommandPointsCostPerAcrion { // пока остались очки на хотя бы одно большое действие
		for bot.doNextMinorCommand == true &&
			bot.minorCommandCount < maxMinorCommandsPerMajorCommand { // делать маленькие действия
			bot.minorCommandCount++
			bot.doCommand()
		}
		bot.majorCommandPointsLeft -= majorCommandPointsCostPerAcrion
		bot.minorCommandCount = 0
		bot.doNextMinorCommand = true
	}

	// botWorld.setBotOnCoord(Coord{bot.CoordX + 1, bot.CoordY + 1}, *bot)
	// WriteLog(fmt.Sprint("Tick ", ThisTickIndex, ": bot ", bot.index, " ", "is on tile ", bot.coordX, ".", bot.coordY, ". Bot energy: ", bot.energy), 4)

	if bot.energy <= 0 {
		botWorld.BotIsDead(bot)
	}
}

// Take value from next genome byte and make a direction out of it (default direction count is 8 so value from 0 to 7, where 0 is up, 1 is up-right and 7 is up-left)
func (bot *Bot) getDirection() int {
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
