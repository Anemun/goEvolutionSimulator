package main

// world config
var worldSizeX = 400
var worldSizeY = 500
var directionsCount byte = 8

// bot config
var botGenomeSize = 64
var organGenomeSize = 8
var maxMinorCommandsPerMajorCommand = 10
var maxBotEnergy = 64
var photosyntesisEnergyGain = 3
var makeChildIfEnergySurplus = true
var mutateChance = 0.25
var maxRelativesDifference = 1

// serialize config
var serializeTickCap = 1000

var filePath = "data/" // linux

//var filePath = "E:/test/"		// windows
var fileNameBase = "chunk_"

// other
var logLevel byte = 2
var maxTickLimit = 2000
var serializeGZ = false

//debug
var debugfillEntireWorld = true
var debugcheckCollisions = false
