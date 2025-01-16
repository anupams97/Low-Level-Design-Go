package firstexample

type VegMania struct {
}

func (p *VegMania) GetPrice() int {
	return 100
}

func (p *VegMania) String() string {
	return "VegMania Pizza"
}
