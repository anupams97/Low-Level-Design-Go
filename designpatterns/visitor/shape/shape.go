package shape

type Shape interface {
	accept(Visitor)
}
