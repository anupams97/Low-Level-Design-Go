package concurrentmarket

type Observer interface {
	OnNotify(event Event)
}
