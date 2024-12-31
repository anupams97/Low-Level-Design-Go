package shape

type Visitor interface {
	visitCircle(*Circle)
	visitSquare(*Square)
}
