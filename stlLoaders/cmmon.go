package stlLoaders

import (
	"os"
	"log"
)

const HeaderSize = 80
const ColorSize = 2
const ChunkSize = 48

type Vector3D struct {
	X float32
	Y float32
	Z float32
}


func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	read, err := file.Read(bytes)
	if err != nil || read != number {
		log.Fatal(err)
	}

	return bytes
}
