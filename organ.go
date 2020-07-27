package main

// Organ class
type Organ struct {
	parent         *Bot
	index          uint64
	coordX         int
	coordY         int
	genome         []byte
	commandPointer int
}

// InitOrgan initializes new organ
func (organ *Organ) InitOrgan(index uint64, parent *Bot, genome []byte) {
	organ.SetCommandPointer(0)
	parent.organs = append(parent.organs, organ)
	organ.parent = parent
	organ.index = index
	organ.genome = genome
}

func (organ *Organ) tick() {

}
