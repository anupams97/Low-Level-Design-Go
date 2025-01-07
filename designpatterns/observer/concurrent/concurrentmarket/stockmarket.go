package concurrentmarket

import (
	"errors"
	"fmt"
	"sync"
)

type StockMarket struct {
	Traders map[Observer]chan Event
	mu      sync.Mutex
	Wg      sync.WaitGroup
}

func NewStockMarket() *StockMarket {
	return &StockMarket{Traders: make(map[Observer]chan Event)}
}
func (s *StockMarket) Register(observer Observer) error {
	s.mu.Lock()

	defer s.mu.Unlock()
	if _, exists := s.Traders[observer]; !exists {
		s.Wg.Add(1)
		s.Traders[observer] = make(chan Event)
		go s.startListening(observer)
		return nil
	}
	return errors.New("trader already registered")
}

func (s *StockMarket) startListening(o Observer) {
	defer s.Wg.Done()
	ch := s.Traders[o]
	for event := range ch {
		o.OnNotify(event)
	}
	fmt.Printf("Stopped listening for trader: %v\n", o)
}

func (s *StockMarket) DeRegister(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	trader, exists := s.Traders[observer]
	if exists {
		close(trader)
		delete(s.Traders, observer)
		s.Wg.Done()
	}
}

func (s *StockMarket) Notify(event Event) {
	for _, ch := range s.Traders {
		ch <- event
	}
}
