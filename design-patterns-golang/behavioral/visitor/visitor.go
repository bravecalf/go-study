package visitor

type Visitor interface {
	visitorForSquare(*Square)
	visitorForCircle(*Circle)
}
