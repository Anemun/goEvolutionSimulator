package main

// command 0
func (bot *Bot) commandSTAY() {
	// WriteLog(fmt.Sprint("Bot ", bot.index, " is doing nothing (command [0]STAY)"), 4)
	bot.IncrementCommandPointer(1)
	bot.doNextMinorCommand = false
}

// command 5
// Look to direction, inrelative to current direction
func (bot *Bot) commandLOOKa() {
	var direction = bot.getDirection()
	var coordToLook = bot.getAdjascentCoordByDirection(direction)
	var objectOnCoord = botWorld.WhatIsOnCoord(coordToLook, bot)
	switch objectOnCoord {
	case "empty":
		bot.IncrementCommandPointer(1)
	case "bot":
		bot.IncrementCommandPointer(2)
	case "relative":
		bot.IncrementCommandPointer(3)
	case "food":
		bot.IncrementCommandPointer(4)
	default:
		panic("There must be one of the values above!")
	}
	bot.doNextMinorCommand = true
	// WriteLog(fmt.Sprint("Bot ", bot.index, " is looking to ", coordToLook, " seeing ", objectOnCoord, " (command [5]LOOKa)"), 4)
}

// command 10
// Move to direction, inrelative to current direction
func (bot *Bot) commandMOVEa() {
	var direction = bot.getDirection()
	var coordToMove = bot.getAdjascentCoordByDirection(direction)
	var objectOnCoord = botWorld.WhatIsOnCoord(coordToMove, nil)
	if objectOnCoord == "empty" { // Также необходимо проверка возможности сдвига всех органов в нужном направлении
		botWorld.setBotOnCoord(coordToMove, bot)
		bot.IncrementCommandPointer(3)
		// WriteLog(fmt.Sprint("Bot ", bot.index, " is moving to", coordToMove, " (command [10]MOVEa)"), 4)
	} else {
		bot.IncrementCommandPointer(2)
		// WriteLog(fmt.Sprint("Bot ", bot.index, " is trying to move to", coordToMove, " but there is ", objectOnCoord, " there (command [10]MOVEa)"), 4)
	}
	// Сдвиг всех органов в нужном направлении

	bot.doNextMinorCommand = false
}

// command 15
func (bot *Bot) commandEAT() {
	// WriteLog(fmt.Sprint("Bot ", bot.index, " NOT IMPLEMENTED", " (command [15]commandEAT)"), 4)
	bot.forwardPointer()
}

// command 20
func (bot *Bot) commandPHOTOSYNTESIS() {
	bot.SetEnergy(photosyntesisEnergyGain)
	// WriteLog(fmt.Sprint("Bot ", bot.index, " gain ", photosyntesisEnergyGain, " energy from photosyntesis", " (command [20]commandPHOTOSYNTESIS)"), 4)
	bot.IncrementCommandPointer(1)
}

// command 25
func (bot *Bot) commandORGAN() {
	// WriteLog(fmt.Sprint("Bot ", bot.index, " NOT IMPLEMENTED", " (command [25]commandORGAN)"), 4)
	bot.forwardPointer()
}

// command 30
func (bot *Bot) commandCHILD() {
	direction := 0
	for direction < int(directionsCount) {
		childCoord := bot.getAdjascentCoordByDirection(direction)
		if botWorld.WhatIsOnCoord(childCoord, nil) != "empty" {
			direction++
		} else {
			botWorld.NewBot(childCoord, bot)
			bot.energy = bot.energy / 2
			// WriteLog(fmt.Sprint("Bot ", bot.index, " create child at ", childCoord, " (command [30]commandCHILD)"), 4)
			bot.IncrementCommandPointer(2)
			return
		}
	}
	// WriteLog(fmt.Sprint("Bot ", bot.index, " tried to create child but there is no room around", " (command [30]commandCHILD)"), 4)
	bot.IncrementCommandPointer(1)
}

func (bot *Bot) forwardPointer() {
	// WriteLog(fmt.Sprint("Bot ", bot.index, " is forwarding pointer from ", bot.CommandPointer(), " to ", bot.CommandPointer()+bot.genome[bot.commandPointer]), 4)
	bot.SetCommandPointer(bot.CommandPointer() + int(bot.genome[bot.commandPointer]))
	bot.doNextMinorCommand = true
}