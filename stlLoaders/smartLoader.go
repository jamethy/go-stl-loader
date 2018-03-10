package stlLoaders

import (
	"log"
	"encoding/binary"
)

type SmartFace struct {
	Normal Vector3D
	A      int
	B      int
	C      int
}

type SmartStlFile struct {
	Header    string
	FaceCount uint32
	Vertices  []Vector3D
	Faces     []SmartFace
}

func SmartRead(path string) (stlFile SmartStlFile) {

	file := mustOpenFile(path)
	defer file.Close()

	stlFile.Header = readHeader(file)
	stlFile.FaceCount = readFaceCount(file)

	stlFile.Faces = make([]SmartFace, stlFile.FaceCount)
	for i := uint32(0); i < stlFile.FaceCount; i++ {
		buffer := mustReadNextBytes(file, ChunkSize)

		if err := binary.Read(buffer, binary.LittleEndian, &stlFile.Faces[i]); err != nil {
			log.Fatal("error while parsing face", err)
		}

		mustReadNextBytes(file, ColorSize)
	}

	return stlFile
}
