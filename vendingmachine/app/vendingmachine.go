package app

import "sync"

var (
	once       sync.Once
	vmInstance *VendingMachine
)

type VendingMachine struct {
	Inventory       *Inventory
	state           State
	amount          int
	selectedProduct *Product
}

func (m *VendingMachine) setState(state State) {
	m.state = state
}

func (m *VendingMachine) resetPayment() {
	m.amount = 0
}

func (m *VendingMachine) resetProduct() {
	m.selectedProduct = nil
}

func NewVendingMachine() *VendingMachine {
	once.Do(func() {
		vmInstance = &VendingMachine{}
		vmInstance.Inventory = NewInventory()
		vmInstance.setState(&Idle{vmInstance})
	})
	return vmInstance
}

func (m *VendingMachine) SelectProduct(pro *Product) {
	m.state.SelectProduct(pro)
}
func (m *VendingMachine) AddCoin(coin Coin) {
	m.state.AddCoin(coin)
}
func (m *VendingMachine) Dispense() {
	m.state.DispenseProduct()
}
func (m *VendingMachine) ReturnChange() {
	m.state.ReturnChange()
}
