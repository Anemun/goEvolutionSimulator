package main

// Define world variable types
type world struct {
	bots   [][]*Bot
	organs [][]*Organ
	food   [][]*food
}

type coordinates struct {
	x int
	y int
}

var nextBotIndex uint64
var thisTickIndex uint64
var aliveBotCount uint64

// Call this to init new world
func (world *world) Init() {

	world.bots = make([][]*Bot, worldSizeX)
	for i := 0; i < worldSizeX; i++ {
		world.bots[i] = make([]*Bot, worldSizeY)
	}

	world.food = make([][]*food, worldSizeX)
	for i := 0; i < worldSizeX; i++ {
		world.food[i] = make([]*food, worldSizeY)
	}

	world.organs = make([][]*Organ, worldSizeX)
	for i := 0; i < worldSizeX; i++ {
		world.organs[i] = make([]*Organ, worldSizeY)
	}

	nextBotIndex = 0
	thisTickIndex = 0
}

func (world *world) Tick() {

	// Cycle through not null (nil!) bots in world coordinates array
	var botList []*Bot
	for i := range world.bots {
		for j := range world.bots[i] {
			if world.bots[i][j] != nil {
				botList = append(botList, world.bots[i][j])
			}
		}
	}
	aliveBotCount = uint64(len(botList))
	for i := range botList {
		botList[i].Tick()
	}

	if serializationEnabled == true {
		serializeTick(thisTickIndex, world.bots, world.food)
	}
	thisTickIndex++

	//   if len(botList) == 0 {
	//     state = "Concluded"
	//   }
}

func (world *world) WhatIsOnCoord(coord coordinates, whoIsAsking *Bot) string {
	coord = world.loopCoords(coord)
	if world.bots[coord.x][coord.y] != nil {
		if whoIsAsking != nil {
			// var whoIsAskingCoord = world.loopCoords(coordinates{whoIsAsking.coordX, whoIsAsking.coordY})
			if world.bots[coord.x][coord.y].index == whoIsAsking.index {
				return "self"
			}
			if world.compareGenome(world.bots[coord.x][coord.y], world.bots[whoIsAsking.coordX][whoIsAsking.coordY]) == true {
				return "relative"
			}
		}
		return "bot"
	}
	if world.food[coord.x][coord.y] != nil {
		return "food"
	}
	if world.organs[coord.x][coord.y] != nil {
		if whoIsAsking != nil {
			if world.organs[coord.x][coord.y].parent == whoIsAsking {
				return "self"
			}
			// var whoIsAskingCoord = world.loopCoords(coordinates{whoIsAsking.coordX, whoIsAsking.coordY})
			if world.compareGenome(world.organs[coord.x][coord.y].parent, world.bots[whoIsAsking.coordX][whoIsAsking.coordY]) == true {
				return "relative"
			}
		}
		return "bot"
	}

	return "empty"
}

func (world *world) BiteObject(coord coordinates, consumer *Bot) {
	if world.food[coord.x][coord.y] != nil {
		world.food[coord.x][coord.y] = nil
		consumer.AddEnergy(foodEnergyGain)
		return
	}

	var victim *Bot = nil
	if world.organs[coord.x][coord.y] != nil {
		victim = world.organs[coord.x][coord.y].parent
	} else if world.bots[coord.x][coord.y] != nil {
		victim = world.bots[coord.x][coord.y]
	}

	if victim != nil {
		var energyAmount = attackEnergyGain
		if victim.energy < energyAmount {
			energyAmount = victim.energy
		}
		victim.AddEnergy(-1 * energyAmount)
		consumer.AddEnergy(energyAmount)
		if victim.energy <= 0 {
			world.BotIsDead(victim)
		}
		return
	}

	if victim == nil {
		return
	}
}

func (world *world) compareGenome(bot1 *Bot, bot2 *Bot) bool {
	var areRelatives bool
	var differenceCount = 0

	if bot1 == nil || bot2 == nil {
		panic("One of the bots are nil! why?")
	}

	for i := range bot1.genome {
		if bot1.genome[i] != bot2.genome[i] {
			differenceCount++
			if differenceCount > maxRelativesDifference {
				areRelatives = false
				return areRelatives
			}
		}
	}
	areRelatives = true
	return areRelatives
}

func (world *world) setBotOnCoord(coord coordinates, bot *Bot) {
	world.bots[bot.coordX][bot.coordY] = nil
	coord = world.loopCoords(coord)
	bot.coordX = coord.x
	bot.coordY = coord.y
	world.bots[bot.coordX][bot.coordY] = bot
}

func (world *world) setOrganOnCoord(coord coordinates, organ *Organ) {
	world.organs[organ.coordX][organ.coordY] = nil
	coord = world.loopCoords(coord)
	organ.coordX = coord.x
	organ.coordY = coord.y
	world.organs[organ.coordX][organ.coordY] = organ
}

func (world *world) NewBot(coord coordinates, parent *Bot) {
	var newBot Bot
	newBot.InitBot(nextBotIndex, parent)
	nextBotIndex++
	world.setBotOnCoord(coord, &newBot)
	// WriteLog(fmt.Sprint("New bot: bot ", newBot.index, " ", "is on tile ", newBot.coordX, ".", newBot.coordY, ". Bot pointer: [", newBot.commandPointer, "]: ", newBot.genome[newBot.commandPointer]), 4)
}

func (world *world) NewOrgan(coord coordinates, parent *Bot, genome []byte) {
	var newOrgan Organ
	newOrgan.InitOrgan(nextBotIndex, parent, genome)
	nextBotIndex++
	world.setOrganOnCoord(coord, &newOrgan)
}

func (world *world) BotIsDead(bot *Bot) {
	world.bots[bot.coordX][bot.coordY].isDead = true
	world.bots[bot.coordX][bot.coordY] = nil
	var foodCoord coordinates
	for i := range bot.organs {
		world.organs[bot.organs[i].coordX][bot.organs[i].coordY] = nil
		foodCoord = coordinates{bot.organs[i].coordX, bot.organs[i].coordY}
		world.food[bot.organs[i].coordX][bot.organs[i].coordY] = &food{foodCoord}
	}

	foodCoord = coordinates{bot.coordX, bot.coordY}
	world.food[bot.coordX][bot.coordY] = &food{foodCoord}
}

func (world *world) GetCurrentTickIndex() uint64 {
	return thisTickIndex
}

func (world *world) loopCoords(coord coordinates) coordinates {
	coord.x = LoopValue(coord.x, 0, worldSizeX)
	coord.y = LoopValue(coord.y, 0, worldSizeY)

	return coord
}
