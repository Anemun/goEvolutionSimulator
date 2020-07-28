package main

// command 0
func (organ *Organ) commandSTAY() {	
	organ.IncrementCommandPointer(1)
	organ.doNextMinorCommand = false
}

// command 20
func (organ *Organ) commandPHOTOSYNTESIS() {
	parent.AddEnergy(photosyntesisEnergyGain)
	// WriteLog(fmt.Sprint("Bot ", bot.index, " gain ", photosyntesisEnergyGain, " energy from photosyntesis", " (command [20]commandPHOTOSYNTESIS)"), 4)
	organ.IncrementCommandPointer(1)
	organ.doNextMinorCommand = false
}