package stlLoaders

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
	IndexMap  map[Vector3D]int
	Faces     []SmartFace
}

func SmartRead(path string) (stlFile SmartStlFile) {

	file := mustOpenFile(path)
	defer file.Close()

	stlFile.Header = readHeader(file)
	stlFile.FaceCount = readFaceCount(file)

	stlFile.Faces = make([]SmartFace, stlFile.FaceCount)
	stlFile.Vertices = make([]Vector3D, stlFile.FaceCount*2)
	stlFile.IndexMap = make(map[Vector3D]int, stlFile.FaceCount*2)

	for i := uint32(0); i < stlFile.FaceCount; i++ {

		face := stlFile.Faces[i]
		face.Normal = readVector3D(file)
		face.A = stlFile.GetIndexOf(readVector3D(file))
		face.B = stlFile.GetIndexOf(readVector3D(file))
		face.C = stlFile.GetIndexOf(readVector3D(file))

		mustReadNextBytes(file, ColorSize)
	}

	return stlFile
}

func (stlFile *SmartStlFile) GetIndexOf(v Vector3D) int {
	oneIndex := stlFile.IndexMap[v]
	if oneIndex == 0 {
		oneIndex = len(stlFile.Vertices)
		stlFile.Vertices = append(stlFile.Vertices, v)
		stlFile.IndexMap[v] = oneIndex
	}

	return oneIndex - 1
}
