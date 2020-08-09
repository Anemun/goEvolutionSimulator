package main

// world config
var worldSizeX = 50
var worldSizeY = 50
var initialBotCount = (worldSizeX + worldSizeY) / 5
var directionsCount byte = 8

// bot config
var botGenomeSize = 64
var initialMajorCommandPointsPerTurn = 10                // Чтобы не заморачиваться с float, вместо кол-ва команд используем очки действия.
var majorCommandPointsCostPerAcrion = 10
var maxMinorCommandsPerMajorCommand = 10
var maxBaseBotEnergy = 64
var botEnergyTickCost = 1
var photosyntesisEnergyGain = 3
var makeChildIfEnergySurplus = true
var childEnergyFraction = 0.5 // how much energy parent gives to child (0.5 - half of current energy)
var mutateChance = 0.25
var maxRelativesDifference = 1
var attackEnergyGain = 5
var foodEnergyGain = 15

// bot organ config
var createOrganCost = 4
var organGenomeSize = 8
var maxOrganMajorCommandsPerTurn = 1
var maxOrganMinorCommandsPerMajorCommand = 3
var maxBaseBotEnergyPerOrgan = 8
var botEnergyTickCostPerOrgan = 1
var organMoveCommandBonusPoints = 5
var eatOrganEnergyGain = 3

// serialize config
var serializationEnabled = false
var serializeTickCap = 100
var filePath = "/home/anemun/evosimData" // linux
//var filePath = "E:/test/"		// windows
var serializeFolderPerDateCap = 100
var fileNameBase = "chunk_"
var serializeGZ = false

// other
var logLevel byte = 5
var maxTickLimit = 2000

//debug
var debugfillEntireWorld = false
var debugcheckCollisions = true
