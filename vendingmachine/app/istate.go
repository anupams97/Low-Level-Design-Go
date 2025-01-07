package app

import "fmt"

type State interface {
	SelectProduct(product *Product)
	AddCoin(coin Coin)
	DispenseProduct()
	ReturnChange()
}

type Idle struct {
	vm *VendingMachine
}

func (i *Idle) SelectProduct(product *Product) {
	if i.vm.Inventory.IsAvailable(product) {
		i.vm.selectedProduct = product
		i.vm.setState(&ReadyState{vm: i.vm})
	} else {
		fmt.Println("item finished")
	}

}

func (i *Idle) AddCoin(coin Coin) {
	fmt.Println("Please select product first")
}

func (i *Idle) AddNote() {
	fmt.Println("Please select product first")
}

func (i *Idle) DispenseProduct() {
	fmt.Println("Please select product first")
}

func (i *Idle) ReturnChange() {
	fmt.Println("Please select product first")
}

type ReadyState struct {
	vm *VendingMachine
}

func (r *ReadyState) SelectProduct(product *Product) {
	fmt.Println("product already selected")
}

func (r *ReadyState) AddCoin(coin Coin) {
	r.vm.amount += int(coin)
	r.checkPaymentStatus()
}

func (r *ReadyState) checkPaymentStatus() {
	if r.vm.amount >= r.vm.selectedProduct.Price {
		r.vm.setState(&DispenseProductState{vm: r.vm})
	}
}

func (r *ReadyState) DispenseProduct() {
	fmt.Println("once money is suff product will be dispensed")
}

func (r *ReadyState) ReturnChange() {
	fmt.Println("after product is dispersed change will be given")
}

type DispenseProductState struct {
	vm *VendingMachine
}

func (d *DispenseProductState) SelectProduct(product *Product) {
	fmt.Println("product is dispering select product again after it is dispersed ")
}

func (d *DispenseProductState) AddCoin(coin Coin) {
	fmt.Println("no need to add more coin for this product")
}

func (d *DispenseProductState) DispenseProduct() {

	fmt.Println("dispensing: ", d.vm.selectedProduct)
	d.vm.Inventory.RemoveProduct(d.vm.selectedProduct)
	d.vm.setState(&ReturnChangeState{d.vm})

}

func (d *DispenseProductState) ReturnChange() {
	fmt.Println("Product is dispersing wait")
}

type ReturnChangeState struct {
	vm *VendingMachine
}

func (r *ReturnChangeState) SelectProduct(product *Product) {
	fmt.Println("Product already dispersed")
}

func (r *ReturnChangeState) AddCoin(coin Coin) {
	fmt.Println("Product already dispersed")
}

func (r *ReturnChangeState) DispenseProduct() {
	fmt.Println("Product already dispersed")
}

func (r *ReturnChangeState) ReturnChange() {
	if r.vm.amount > r.vm.selectedProduct.Price {
		fmt.Println("refund ", r.vm.amount-r.vm.selectedProduct.Price)

	}
	r.vm.resetPayment()
	r.vm.resetProduct()
	r.vm.setState(&Idle{r.vm})
}
