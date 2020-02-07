package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"github.com/vmihailenco/msgpack"
)

// Каждый тик упаковываем ботов в мир
// Кладём в хранилку тиков
// Когда кол-во тиков становится равно пороговому значению - сериализуем

type ChunkPack struct {
	WorldSizeX int
	WorldSizeY int
	ChunkIndex int
	Ticks      []TickPack
}

func (chunk *ChunkPack) Reset() {
	chunk.ChunkIndex = 0
	chunk.Ticks = nil
}

type TickPack struct {
	TickIndex uint64
	Bots      []BotPack
}

type BotPack struct {
	//Index uint64
	//Genome string
	CoordX byte
	CoordY byte
	Energy byte
	// Organs []*organ
}

var currentChunk ChunkPack
var nextChunkIndex int

var serializationFolderPath string

// Message Pack

// SerializeTick puts current tick to chunk, serialize chunk if treshhold is exceeded
func SerializeTick(tickIndex uint64, bots [][]*Bot) {
	// WriteLog(fmt.Sprint("Serializing tick #", tickIndex), 4)

	var botsToPack []BotPack
	for i := range bots {
		for j := range bots[i] {
			if bots[i][j] != nil {
				var worldBot = bots[i][j]
				bp := BotPack{
					//Index:  worldBot.index,
					CoordX: byte(worldBot.coordX),
					CoordY: byte(worldBot.coordY),
					Energy: byte(worldBot.energy)}
				//Genome: worldBot.genome}
				// Organs: worldBot.organs}
				botsToPack = append(botsToPack, bp)
			}
		}
	}

	tickPack := TickPack{
		TickIndex: tickIndex,
		Bots:      botsToPack}

	currentChunk.WorldSizeX = worldSizeX
	currentChunk.WorldSizeY = worldSizeY
	currentChunk.Ticks = append(currentChunk.Ticks, tickPack)
	currentChunk.ChunkIndex = nextChunkIndex

	if len(currentChunk.Ticks) >= serializeTickCap {
		seriazizeChunk(&currentChunk)
		currentChunk.Reset()
		nextChunkIndex++
	}
}

func seriazizeChunk(chunk *ChunkPack) {
	// WriteLog(fmt.Sprint("Serializing chunk #", nextChunkIndex, " (", len(chunk.Ticks), " ticks in chunk)"), 4)

	data, err := msgpack.Marshal(&chunk)
	if err != nil {
		log.Fatal("marshal error: ", err)
	}

	if serializeGZ == true {
		// GZ Serialization
		filename := fmt.Sprint(serializationFolderPath, fileNameBase, nextChunkIndex, ".gz")

		if _, err := os.Stat(serializationFolderPath); os.IsNotExist(err) {
			os.MkdirAll(serializationFolderPath, os.ModePerm)
		}

		// Open a file for writing.
		f, createErr := os.Create(filename)

		if createErr != nil {
			log.Fatal("create error: ", createErr)
		}

		// Create gzip writer.
		w, _ := gzip.NewWriterLevel(f, gzip.BestSpeed)

		// Write bytes in compressed form to the file.
		_, writeErr := w.Write(data)

		// Close the file.
		w.Close()

		if writeErr != nil {
			log.Fatal("write error: ", writeErr)
		} else {
			// WriteLog(fmt.Sprint("File saved as ", filename), 4)
		}
	} else {
		// BIN Sezialization
		filename := fmt.Sprint(serializationFolderPath, fileNameBase, nextChunkIndex, ".bin")

		if _, err := os.Stat(serializationFolderPath); os.IsNotExist(err) {
			os.MkdirAll(serializationFolderPath, os.ModePerm)
		}

		writeErr := ioutil.WriteFile(filename, data, 0644)

		if writeErr != nil {
			log.Fatal("write error: ", writeErr)
		}
	}

	// // Unmarshaling test
	// var deserChunk ChunkPack

	// err = msgpack.Unmarshal(data, &deserChunk)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := range deserChunk.Ticks {
	// 	fmt.Println(deserChunk.Ticks[i])
	// }
}

// FinalSerialization finalize current chunk even if incomplete
func FinalSerialization() {
	seriazizeChunk(&currentChunk)
}

func CreateFolders() {
	todayDate := time.Now().Format("2006-01-02")
	serializationFolderPath = filePath + "/" + todayDate + "/"
	os.RemoveAll(serializationFolderPath)
	os.MkdirAll(serializationFolderPath, os.ModePerm)
}

// func TestMsgPack(bot *bot) {
// 	b, err := msgpack.Marshal(&bot)
// 	if err != nil {
// 		panic(err)
// 	}

// 	filename := fmt.Sprint("data/", "msG", fileNameBase, nextChunkIndex, ".bin")
// 	writeErr := ioutil.WriteFile(filename, b, 0644)
// 	if writeErr != nil {
// 		log.Fatal("write error: ", writeErr)
// 	}

// 	var botNew mT
// 	err = msgpack.Unmarshal(b, &botNew)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(botNew.TestString)
// }

// // Protobuf

// type chunk struct {
// 	ticks []*TickMessage
// }

// type tick struct {
// 	index uint64
// 	bots  []BotMessage
// }

// // SerializeTick Put current tick data into chunk, if chink's size is more than cap, write chunk to file
// func SerializeTick(index uint64, bots [][]*bot) {
// 	var b []*BotMessage

// 	for i := range bots {
// 		for j := range bots[i] {
// 			if bots[i][j] != nil {
// 				bot := bots[i][j]

// 				botMsg := &BotMessage{
// 					BotIndex:  bot.index,
// 					CoordX:    bot.coordX,
// 					CoordY:    bot.coordY,
// 					Energy:    bot.energy,
// 					BotGenome: bot.genome}

// 				b = append(b, botMsg)
// 			}
// 		}
// 	}

// 	tickMsg := &TickMessage{
// 		TickIndex:   index,
// 		BotsInWorld: b}

// 	currentChunk.ticks = append(currentChunk.ticks, tickMsg)

// 	if len(currentChunk.ticks) >= serializeTickCap {
// 		serializeChunk(&currentChunk)
// 		nextChunkIndex++
// 	}
// }

// func serializeChunk(chunk *chunk) {
// 	// Маршалим чанк
// 	chunkMsg := &ChunkMessage{
// 		TicksInChunk: chunk.ticks}
// 	data, marshErr := proto.Marshal(chunkMsg)
// 	if marshErr != nil {
// 		log.Fatal("marshaling error: ", marshErr)
// 		return
// 	}

// 	// Сохраняем как файл
// 	filename := fmt.Sprint("data/", fileNameBase, nextChunkIndex, ".bin")
// 	writeErr := ioutil.WriteFile(filename, data, 0644)
// 	if writeErr != nil {
// 		log.Fatal("write error: ", writeErr)
// 	}

// 	// Обнуляем чанк
// 	chunk.Reset()
// }
