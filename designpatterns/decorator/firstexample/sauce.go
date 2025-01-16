package firstexample

import "fmt"

type Sauce struct {
	Pizza IPizza
}

func (p *Sauce) GetPrice() int {
	return p.Pizza.GetPrice() + 5
}

func (p *Sauce) String() string {
	return fmt.Sprintf("%s + Sauce", p.Pizza)
}

func NewSauce(pizza IPizza) IPizza {
	return &Sauce{Pizza: pizza}
}
