package main

// world config
var worldSizeX = 200
var worldSizeY = 250
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
var serializeTickCap = 100

//var filePath = "data/"
var filePath = "E:/test/"
var fileNameBase = "chunk_"

// other
var logLevel byte = 1
var maxTickLimit = 500
var serializeGZ = true

//debug
var debugfillEntireWorld = false
var debugcheckCollisions = false
