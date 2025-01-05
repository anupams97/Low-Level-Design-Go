package stockmarket

type Notifier interface {
	Register(Observer)
	DeRegister(Observer)
	Notify(event Event)
}
