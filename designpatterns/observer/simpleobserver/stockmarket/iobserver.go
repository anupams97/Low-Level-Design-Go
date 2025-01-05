package stockmarket

type Observer interface {
	OnNotify(event Event)
}
