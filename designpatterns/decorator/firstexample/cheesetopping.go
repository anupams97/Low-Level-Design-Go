package firstexample

import "fmt"

type CheeseTopping struct {
	Pizza IPizza
}

func (p *CheeseTopping) GetPrice() int {
	return p.Pizza.GetPrice() + 20
}

func (p *CheeseTopping) String() string {
	return fmt.Sprintf("%s + Cheese", p.Pizza)
}

func NewCheeseTopping(pizza IPizza) IPizza {
	return &CheeseTopping{Pizza: pizza}
}
