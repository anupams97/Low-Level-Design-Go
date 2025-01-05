package stockmarket

type StockMarket struct {
	Traders map[Observer]struct{}
}

func (s *StockMarket) Register(observer Observer) {
	s.Traders[observer] = struct{}{}
}

func (s *StockMarket) DeRegister(observer Observer) {
	delete(s.Traders, observer)
}

func (s *StockMarket) Notify(event Event) {
	for ob, _ := range s.Traders {
		ob.OnNotify(event)
	}
}
