package stlLoaders

import (
	"os"
	"log"
	"bytes"
	"encoding/binary"
)

type Face struct {
	Normal Vector3D
	A      Vector3D
	B      Vector3D
	C      Vector3D
}

type StlFile struct {
	Header    string
	FaceCount uint32
	Faces     []Face
}

func BasicRead(path string) (stlFile StlFile) {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal("error while reading file", err)
	}
	defer file.Close()

	stlFile.Header = string(readNextBytes(file, HeaderSize))

	data := readNextBytes(file, 4)
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.LittleEndian, &stlFile.FaceCount); err != nil {
		log.Fatal("error while parsing face count", err)
	}

	stlFile.Faces = make([]Face, stlFile.FaceCount)
	for i := uint32(0); i < stlFile.FaceCount; i++ {
		data = readNextBytes(file, ChunkSize)

		buffer = bytes.NewBuffer(data)
		if err = binary.Read(buffer, binary.LittleEndian, &stlFile.Faces[i]); err != nil {
			log.Fatal("error while parsing face", err)
		}

		readNextBytes(file, ColorSize)
	}

	return stlFile
}
