package main

import (
	"math/rand"
)

// Bot variables
type Bot struct {
	index              uint64
	coordX             int
	coordY             int
	energy             int
	genome             []byte
	organs             []*organ
	commandPointer     int
	isDead             bool
	doNextMinorCommand bool
	minorCommandCount  int
	majorCommandLeft   int
}

type organ struct {
	parent *Bot
	coordX int
	coordY int
	genome []int
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

// SetEnergy SET
func (bot *Bot) SetEnergy(increment int) {
	bot.energy += increment
	if bot.energy > maxBotEnergy {
		if makeChildIfEnergySurplus == true {
			bot.commandCHILD()
		} else {
			bot.energy = maxBotEnergy
		}
	}
}

// NewBot creates new bot
func (bot *Bot) NewBot(index uint64, parent *Bot) {
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
		bot.energy = parent.energy / 2
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
	case 10:
		bot.commandMOVEa()
	case 15:
		bot.commandEAT()
	case 20:
		bot.commandPHOTOSYNTESIS()
	case 25:
		bot.commandORGAN()
	case 30:
		bot.commandCHILD()
	default:
		bot.forwardPointer()
	}
}

// Tick bot logic
func (bot *Bot) Tick() {
	if bot.isDead {
		return
	}

	bot.energy--
	bot.minorCommandCount = 0
	bot.doNextMinorCommand = true
	bot.majorCommandLeft = 1

	for i := range bot.organs {
		bot.organs[i].tick()
	}

	// Бот имеет несколько больших действий, в каждом есть энное количество маленьких
	for bot.majorCommandLeft > 0 { // пока осталось хотя бы одно большое действие
		for bot.doNextMinorCommand == true &&
			bot.minorCommandCount < maxMinorCommandsPerMajorCommand { // делать маленькие действия
			bot.minorCommandCount++
			bot.doCommand()
		}
		bot.majorCommandLeft--
		bot.minorCommandCount = 0
		bot.doNextMinorCommand = true
	}

	// botWorld.setBotOnCoord(Coord{bot.CoordX + 1, bot.CoordY + 1}, *bot)
	// WriteLog(fmt.Sprint("Tick ", ThisTickIndex, ": bot ", bot.index, " ", "is on tile ", bot.coordX, ".", bot.coordY, ". Bot energy: ", bot.energy), 4)

	if bot.energy <= 0 {
		botWorld.BotIsDead(bot)
	}
}

func (organ *organ) tick() {

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

	var newCoord = coordinates{bot.coordX, bot.coordY}
	switch direction {
	case 0:
		newCoord.y++
	case 1:
		newCoord.x++
		newCoord.y++
	case 2:
		newCoord.x++
	case 3:
		newCoord.x++
		newCoord.y--
	case 4:
		newCoord.y--
	case 5:
		newCoord.x--
		newCoord.y--
	case 6:
		newCoord.x--
	case 7:
		newCoord.x--
		newCoord.y++
	}
	newCoord.x = LoopValue(newCoord.x, 0, worldSizeX)
	newCoord.y = LoopValue(newCoord.y, 0, worldSizeY)

	return newCoord
}