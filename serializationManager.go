package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

// Каждый тик упаковываем ботов в мир
// Кладём в хранилку тиков
// Когда кол-во тиков становится равно пороговому значению - сериализуем

var currentChunk ChunkMessage
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

// Protobuf
// protoc --proto_path=. --go_out=. data.proto

// serializeTick Put current tick data into chunk, if chink's size is more than cap, write chunk to file
func serializeTick(index uint64, bots [][]*Bot) {
	var b []*BotMessage

	for i := range bots {
		for j := range bots[i] {
			if bots[i][j] != nil {
				bot := bots[i][j]

				botMsg := &BotMessage{
					Index:  bot.index,
					CoordX: uint32(bot.coordX),
					CoordY: uint32(bot.coordY),
					Energy: uint32(bot.energy),
					Genome: bot.genome}

				b = append(b, botMsg)
			}
		}
	}

	tickMsg := &TickMessage{
		TickIndex: index,
		Bots:      b}

	currentChunk.Ticks = append(currentChunk.Ticks, tickMsg)

	if len(currentChunk.Ticks) >= serializeTickCap {
		serializeChunk(&currentChunk)
		nextChunkIndex++
	}
}

func serializeChunk(chunk *ChunkMessage) {
	// Маршалим чанк
	// chunkMsg := &ChunkMessage{
	// 	Ticks: chunk.ticks}
	chunk.ChunkIndex = uint32(nextChunkIndex)
	data, marshErr := proto.Marshal(chunk)
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

	chunk.Ticks = nil
}
