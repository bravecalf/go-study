package visitor

func TestVisitorDemo() {
	square := &Square{side: 2}
	circle := &Circle{radius: 2}

	areaVisitor := &AreaVisitor{}

	square.accept(areaVisitor)
	circle.accept(areaVisitor)
}
