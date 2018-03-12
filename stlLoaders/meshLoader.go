package stlLoaders

type MeshPoint struct {
	Point           Vector3D
	ConnectedPoints []int
	ConnectedFaces  []int
	Index           int
}

type MeshFace struct {
	Face  SmartFace
	Index int
}

type MeshStlFile struct {
	Header        string
	FaceCount     uint32
	Vertices      []MeshPoint
	PointIndexMap map[Vector3D]int
	FaceIndexMap  map[SmartFace]int
	Faces         []MeshFace
}

func MeshRead(path string) (stlFile MeshStlFile) {

	file := mustOpenFile(path)
	defer file.Close()

	stlFile.Header = readHeader(file)
	stlFile.FaceCount = readFaceCount(file)

	stlFile.Faces = make([]MeshFace, stlFile.FaceCount)
	stlFile.PointIndexMap = map[Vector3D]int{}
	stlFile.FaceIndexMap = map[SmartFace]int{}

	for i := uint32(0); i < stlFile.FaceCount; i++ {

		meshFace := stlFile.Faces[i]
		meshFace.Index = int(i)
		face := meshFace.Face
		face.Normal = readVector3D(file)
		face.A = stlFile.GetIndexOf(readVector3D(file), meshFace.Index)
		face.B = stlFile.GetIndexOf(readVector3D(file), meshFace.Index)
		face.C = stlFile.GetIndexOf(readVector3D(file), meshFace.Index)

		mustReadNextBytes(file, ColorSize)
	}

	for _, point := range stlFile.Vertices {
		point.ConnectedFaces = unique(point.ConnectedFaces)
		for _, faceIndex := range point.ConnectedFaces {
			face := stlFile.Faces[faceIndex]
			point.ConnectedPoints = append(point.ConnectedPoints, face.Face.A, face.Face.B, face.Face.C)
		}
		point.ConnectedPoints = unique(point.ConnectedPoints)
	}

	return stlFile
}

func unique(intSlice []int) (list []int) {
	keys := make(map[int]bool)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (stlFile *MeshStlFile) GetIndexOf(v Vector3D, face int) int {

	oneIndex := stlFile.PointIndexMap[v]
	index := oneIndex - 1

	if oneIndex == 0 {
		oneIndex = len(stlFile.Vertices)
		point := MeshPoint{v, []int{}, []int{face}, index}
		stlFile.Vertices = append(stlFile.Vertices, point)
		stlFile.PointIndexMap[v] = oneIndex
	} else {
		point := stlFile.Vertices[index]
		point.ConnectedFaces = append(point.ConnectedFaces, face)
	}
	return index
}
