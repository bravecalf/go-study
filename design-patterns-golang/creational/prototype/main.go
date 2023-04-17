package prototype

func TestProtoTypeDemo() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{
		children: []INode{file1, file2, file3},
		name:     "folder1",
	}

	folder2 := &Folder{
		children: []INode{folder1, file2, file3},
		name:     "folder2",
	}
	folder2.print(" ")

	copyFolder2 := folder2.clone()
	copyFolder2.print(" ")

}
