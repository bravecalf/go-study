package composite

func TestCompositeDemo() {
	f1 := &File{"file1"}
	f2 := &File{"file2"}
	f1.search("okk")

	fd1 := &Folder{name: "folder1"}
	fd1.addComponent(f1)
	fd1.addComponent(f2)
	fd1.search("hah")
}
