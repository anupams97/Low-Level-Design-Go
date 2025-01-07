package app

type ProductType string

const (
	Snack ProductType = "Chips"
	Drink ProductType = "Drink"
	Milk  ProductType = "Milk"
)

type Product struct {
	ProductType ProductType
	Price       int
}
