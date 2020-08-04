package main

// command 0
func (organ *Organ) commandSTAY() {
	organ.IncrementCommandPointer(1)
	organ.doNextMinorCommand = false
}

// command 20
func (organ *Organ) commandPHOTOSYNTESIS() {
	organ.parent.AddEnergy(photosyntesisEnergyGain)
	// WriteLog(fmt.Sprint("Bot ", bot.index, " gain ", photosyntesisEnergyGain, " energy from photosyntesis", " (command [20]commandPHOTOSYNTESIS)"), 4)
	organ.IncrementCommandPointer(1)
	organ.doNextMinorCommand = false
}

func (organ *Organ) forwardPointer() {
	// WriteLog(fmt.Sprint("Bot ", bot.index, " is forwarding pointer from ", bot.CommandPointer(), " to ", bot.CommandPointer()+bot.genome[bot.commandPointer]), 4)
	organ.SetCommandPointer(organ.CommandPointer() + int(organ.genome[organ.commandPointer]))
	organ.doNextMinorCommand = true
}
