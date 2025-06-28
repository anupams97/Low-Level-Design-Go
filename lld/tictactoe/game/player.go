package game

type Player struct {
	name string
	sign Sign
}

func (p *Player) GetName() string {
	return p.name
}
func (p *Player) GetSign() Sign {
	return p.sign
}

func NewPlayer(name string, sign Sign) *Player {
	return &Player{name: name, sign: sign}
}
