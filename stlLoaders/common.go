package stlLoaders

import (
	"os"
	"log"
	"bytes"
	"encoding/binary"
)

const HeaderSize = 80
const ColorSize = 2
const ChunkSize = 48

type Vector3D struct {
	X float32
	Y float32
	Z float32
}

func readHeader(file *os.File) string {
	return mustReadNextBytes(file, HeaderSize).String()
}

func readFaceCount(file *os.File) (faceCount uint32) {

	buffer := mustReadNextBytes(file, 4)
	if err := binary.Read(buffer, binary.LittleEndian, &faceCount); err != nil {
		log.Fatal("error while parsing face count", err)
	}

	return faceCount
}

func readVector3D(file *os.File) (vec Vector3D) {

	buffer := mustReadNextBytes(file, 12)

	if err := binary.Read(buffer, binary.LittleEndian, &vec); err != nil {
		log.Fatal("error while parsing vector", err)
	}

	return vec
}

func mustReadNextBytes(file *os.File, number int) *bytes.Buffer {
	data := make([]byte, number)

	read, err := file.Read(data)
	if err != nil || read != number {
		log.Fatal(err)
	}

	return bytes.NewBuffer(data)
}

func mustOpenFile(path string) *os.File {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal("error while reading file", err)
	}

	return file
}
