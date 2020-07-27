package main

// world config
var worldSizeX = 400
var worldSizeY = 500
var initialBotCount = (worldSizeX + worldSizeY) / 5
var directionsCount byte = 8

// bot config
var botGenomeSize = 64
var organGenomeSize = 8
var maxMajorCommandsPerTurn = 1
var maxMinorCommandsPerMajorCommand = 10
var maxBaseBotEnergy = 64
var maxBaseBotEnergyPerOrgan = 8
var botEnergyTickCost = 1
var botEnergyTickCostPerOrgan = 1
var photosyntesisEnergyGain = 3
var makeChildIfEnergySurplus = true
var childEnergyFraction = 0.5 // how much energy parent gives to child (0.5 - half of current energy)
var createOrganCost = 4
var mutateChance = 0.25
var maxRelativesDifference = 1

// serialize config
var serializeTickCap = 100
var filePath = "/home/anemun/evosimData" // linux
//var filePath = "E:/test/"		// windows
var serializeFolderPerDateCap = 100
var fileNameBase = "chunk_"

// other
var logLevel byte = 5
var maxTickLimit = 2000
var serializeGZ = false

//debug
var debugfillEntireWorld = false
var debugcheckCollisions = true
