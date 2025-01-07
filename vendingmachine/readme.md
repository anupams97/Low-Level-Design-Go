Requirements:
1. We have vending machine with products in it.
2. user can select a product.
3. after that add the money into the machine
4. if money sufficient dispenses the  product 
5. returns the change if any.
6. user can cancel the process at any time if the product is not dispensed, once the product is dispensed
    process cannot be canceled.

Entities:
Vending Machine 
    Inventory
        Products

Coin
Note

State:
    1.Idle State
    2.Select Product State
    3.Add Money State
    4.Dispense Product state
    5.Return change state
    6.Cancel state