package main

//go:generate msgp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// Каждый тик упаковываем ботов в мир
// Кладём в хранилку тиков
// Когда кол-во тиков становится равно пороговому значению - сериализуем

var currentChunk SerChunk
var nextChunkIndex int

var serializationFolderPath string

// FinalSerialization finalize current chunk even if incomplete
func FinalSerialization() {
	serializeChunk(&currentChunk)
}

func createFolders() {
	todayDate := time.Now().Format("2006-01-02")
	var folderIndex = 0

	for folderIndex < serializeFolderPerDateCap {
		serializationFolderPath = filePath + "/" + todayDate + "_" + strconv.Itoa(folderIndex) + "/"
		if _, err := os.Stat(serializationFolderPath); os.IsNotExist(err) {
			os.MkdirAll(serializationFolderPath, os.ModePerm)
			break
		}
		folderIndex++
	}
}

type SerBot struct {
	BotIndex  uint64
	BotCoordX int
	BotCoordY int
	BotEnergy int
	BotGenome []byte
	BotOrgans []*SerOrgan
}
type SerOrgan struct {
	ParentBotIndex uint64
	OrganIndex     uint64
	OrganCoordX    int
	OrganCoordY    int
	OrganEnergy    int
	OrganGenome    []byte
}

type SerTick struct {
	TickIndex uint64
	Bots      []*SerBot
}
type SerChunk struct {
	ChunkIndex uint
	WorldSizeX uint
	WorldSizeY uint
	Ticks      []*SerTick
}

func (b *SerBot) preprareBot(index uint64, x int, y int) *SerBot {
	b.BotIndex = index
	b.BotCoordX = x
	b.BotCoordY = y

	return b
}

func serializeTick(tickIndex uint64, bots [][]*Bot) {
	var tickBots []*SerBot
	for i := range bots {
		for j := range bots[i] {
			if bots[i][j] != nil {
				bot := bots[i][j]

				botToSerz := &SerBot{
					BotIndex:  bot.index,
					BotCoordX: bot.coordX,
					BotCoordY: bot.coordY,
					BotEnergy: bot.energy,
					BotGenome: bot.genome}

				for o := range bots[i][j].organs {
					if bots[i][j].organs[o] != nil {
						organ := bots[i][j].organs[o]

						organToSerz := &SerOrgan{
							ParentBotIndex: bot.index,
							OrganIndex:     organ.index,
							OrganCoordX:    organ.coordX,
							OrganCoordY:    organ.coordY,
							OrganGenome:    organ.genome}
						botToSerz.BotOrgans = append(botToSerz.BotOrgans, organToSerz)
					}
				}
				tickBots = append(tickBots, botToSerz)

				// var newBot SerBot
				// newTick.Bots = append(newTick.Bots, newBot.preprareBot(bots[i][j].index, bots[i][j].coordX, bots[i][j].coordY))
			}
		}
	}
	// data, _ := newTick.MarshalMsg(nil)
	// print(data)
	tickSerz := &SerTick{
		TickIndex: tickIndex,
		Bots:      tickBots,
	}
	currentChunk.Ticks = append(currentChunk.Ticks, tickSerz)
	if len(currentChunk.Ticks) >= serializeTickCap {
		serializeChunk(&currentChunk)
		nextChunkIndex++
	}
}

func serializeChunk(chunk *SerChunk) {
	// Маршалим чанк
	// chunkMsg := &ChunkMessage{
	// 	Ticks: chunk.ticks}
	chunk.ChunkIndex = uint(nextChunkIndex)
	chunk.WorldSizeX = uint(worldSizeX)
	chunk.WorldSizeY = uint(worldSizeY)

	data, marshErr := chunk.MarshalMsg(nil)
	if marshErr != nil {
		log.Fatal("marshaling error: ", marshErr)
		return
	}

	// Сохраняем как файл
	filename := fmt.Sprint(serializationFolderPath, fileNameBase, nextChunkIndex, ".bin")
	writeErr := ioutil.WriteFile(filename, data, 0644)
	if writeErr != nil {
		log.Fatal("write error: ", writeErr)
	}
}

// // Protobuf
// // protoc --proto_path=. --go_out=. data.proto

// // serializeTick Put current tick data into chunk, if chink's size is more than cap, write chunk to file
// func serializeTick(index uint64, bots [][]*Bot) {
// 	var b []*BotMessage

// 	for i := range bots {
// 		for j := range bots[i] {
// 			if bots[i][j] != nil {
// 				bot := bots[i][j]

// 				botMsg := &BotMessage{
// 					Index:  bot.index,
// 					CoordX: uint32(bot.coordX),
// 					CoordY: uint32(bot.coordY),
// 					Energy: uint32(bot.energy),
// 					Genome: bot.genome}

// 				for o := range bots[i][j].organs {
// 					if bots[i][j].organs[o] != nil {
// 						organ := bots[i][j].organs[o]

// 						organMsg := &OrganMessage{
// 							ParentBotIndex: bot.index,
// 							OrganIndex:     organ.index,
// 							CoordX:         uint32(organ.coordX),
// 							CoordY:         uint32(organ.coordY),
// 							Genome:         organ.genome}
// 						botMsg.Organs = append(botMsg.Organs, organMsg)
// 					}
// 				}
// 				b = append(b, botMsg)
// 			}
// 		}
// 	}

// 	tickMsg := &TickMessage{
// 		TickIndex: index,
// 		Bots:      b}

// 	currentChunk.Ticks = append(currentChunk.Ticks, tickMsg)

// 	if len(currentChunk.Ticks) >= serializeTickCap {
// 		serializeChunk(&currentChunk)
// 		nextChunkIndex++
// 	}
// }

// func serializeChunk(chunk *ChunkMessage) {
// 	// Маршалим чанк
// 	// chunkMsg := &ChunkMessage{
// 	// 	Ticks: chunk.ticks}
// 	chunk.ChunkIndex = uint32(nextChunkIndex)
// 	chunk.WorldSizeX = uint32(worldSizeX)
// 	chunk.WorldSizeY = uint32(worldSizeY)
// 	data, marshErr := proto.Marshal(chunk)
// 	if marshErr != nil {
// 		log.Fatal("marshaling error: ", marshErr)
// 		return
// 	}

// 	chunk.Ticks = nil
// }
