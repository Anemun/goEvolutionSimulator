package main

// Organ class
type Organ struct {
	parent             *Bot
	index              uint64
	coordX             int
	coordY             int
	genome             []byte
	commandPointer     int
	doNextMinorCommand bool
	minorCommandCount  int
	majorCommandLeft   int
}

// InitOrgan initializes new organ
func (organ *Organ) InitOrgan(index uint64, parent *Bot, genome []byte) {
	organ.SetCommandPointer(0)
	parent.organs = append(parent.organs, organ)
	organ.parent = parent
	organ.index = index
	organ.genome = genome
}

// CommandPointer GET
func (organ *Organ) CommandPointer() int {
	return organ.commandPointer
}

// SetCommandPointer SET
func (organ *Organ) SetCommandPointer(newPointer int) {
	organ.commandPointer = LoopValue(newPointer, 0, organGenomeSize)
}

// IncrementCommandPointer ++
func (organ *Organ) IncrementCommandPointer(increment int) {
	organ.SetCommandPointer(organ.CommandPointer() + increment)
}

func (organ *Organ) doCommand() {
	//Organ commands defined in organCommands.go
	switch organ.genome[organ.commandPointer] {
	case 0:
		organ.commandSTAY()
	case 20:
		organ.commandPHOTOSYNTESIS()
	default:
		organ.forwardPointer()
	}
}

func (organ *Organ) tick() {
	if organ.parent.isDead {
		return
	}

	organ.minorCommandCount = 0
	organ.doNextMinorCommand = true
	organ.majorCommandLeft = maxOrganMajorCommandsPerTurn

	for organ.majorCommandLeft > 0 { // пока осталось хотя бы одно большое действие
		for organ.doNextMinorCommand == true &&
			organ.minorCommandCount < maxOrganMinorCommandsPerMajorCommand { // делать маленькие действия
			organ.minorCommandCount++
			organ.doCommand()
		}
		organ.majorCommandLeft--
		organ.minorCommandCount = 0
		organ.doNextMinorCommand = true
	}
}

func (organ *Organ) getAdjascentCoordByDirection(direction int) coordinates {
	if directionsCount != 8 {
		panic("This function covers only 8 directions! Update it if you change directions count")
	}

	var targetCoord = coordinates{organ.coordX, organ.coordY}
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
