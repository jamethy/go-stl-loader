package stlLoaders

import (
	"log"
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

	file := mustOpenFile(path)
	defer file.Close()

	stlFile.Header = readHeader(file)
	stlFile.FaceCount = readFaceCount(file)

	stlFile.Faces = make([]Face, stlFile.FaceCount)
	for i := uint32(0); i < stlFile.FaceCount; i++ {

		buffer := mustReadNextBytes(file, ChunkSize)

		if err := binary.Read(buffer, binary.LittleEndian, &stlFile.Faces[i]); err != nil {
			log.Fatal("error while parsing face", err)
		}

		mustReadNextBytes(file, ColorSize)
	}

	return stlFile
}
