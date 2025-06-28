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


URl Shortner

Function Requirements
1. Add long URL in the system, get a shorter URL
2. 0.5 million req/s get
3. 30 Days ret.

Non Functional Requirements
1. Available
2. low latency

Capacity Planning

1. storage long url     200 bytes
2. storage short url

database  shorturl

further - bits  40 bits  5 bytes


0.5 * 10 ^6  * 60* 24 *30  * 240  18000*10^6 bits - 5  TB

Building block
shortnening service
30 days retention period
low latency
NO SQL


Part 1 :
API : /Post shortURL
response
req { long url }

if some longURL  there use the predefined short url only
create :
1. long URL hash -> shorter
db id  short url

how do we create a short url
uniqueID generators
long Url  -> shortURL
DB
SQL
id shortURL LongURL


Part 2:
/Get request shortURL
we can check if the short uRL is present

servers handle

LB -

Cache short - frequently short URLS served the cache itself.
Deciding what type of cache


