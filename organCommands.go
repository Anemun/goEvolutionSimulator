package main

// command 0
func (organ *Organ) commandSTAY() {
	organ.IncrementCommandPointer(1)
	organ.doNextMinorCommand = false
}

// command 10
func (organ *Organ) commandSPEEDUP() {
	organ.IncrementCommandPointer(1)
	organ.parent.majorCommandPointsLeft += organMoveCommandBonusPoints
	organ.doNextMinorCommand = false
}

func (organ *Organ) commandEAT() {
  var direction = organ.getDirection()
  var coordToBite = organ.getAdjascentCoordByDirection(direction)
	var objectOnCoord = botWorld.WhatIsOnCoord(coordToBite, nil)
  switch objectOnCoord {
	case "empty":
		organ.IncrementCommandPointer(2)
	case "bot":
		organ.IncrementCommandPointer(3)
	case "relative":
		organ.IncrementCommandPointer(3)        // так специально, своих и чужих не различаем, их должна различить команда ПОСМОТРЕТЬ бота #TODO подумать об этом, орган не может смотреть
	case "food":
		organ.IncrementCommandPointer(4)
	case "self":
		organ.IncrementCommandPointer(5)
	default:
		panic("There must be one of the values above!")
	}
  botWorld.BiteObject(coordToBite, organ.parent)
	// WriteLog(fmt.Sprint("Organ ", organ.index, " bites ", objectOnCoord, " bot gains energy, " (command [15]commandEAT)"), 4)
	organ.doNextMinorCommand = false
}

// command 20
func (organ *Organ) commandPHOTOSYNTESIS() {
	organ.parent.AddEnergy(photosyntesisEnergyGain)
	// WriteLog(fmt.Sprint("Bot ", bot.index, " gain ", photosyntesisEnergyGain, " energy from organ photosyntesis", " (command [20]commandPHOTOSYNTESIS)"), 4)
	organ.IncrementCommandPointer(1)
	organ.doNextMinorCommand = false
}

func (organ *Organ) forwardPointer() {
	// WriteLog(fmt.Sprint("Bot ", bot.index, " is forwarding pointer from ", bot.CommandPointer(), " to ", bot.CommandPointer()+bot.genome[bot.commandPointer]), 4)
	organ.SetCommandPointer(organ.CommandPointer() + int(organ.genome[organ.commandPointer]))
	organ.doNextMinorCommand = true
}
