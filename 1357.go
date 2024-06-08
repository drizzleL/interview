package main

type Cashier struct {
	priceDict map[int]int
	discount  float64
	lucky     int
	customers int
}

func CashConstructor(n int, discount int, products []int, prices []int) Cashier {
	priceDict := map[int]int{}
	for i := 0; i < len(prices); i++ {
		priceDict[products[i]] = prices[i]
	}
	return Cashier{
		lucky:     n,
		discount:  float64(discount),
		priceDict: priceDict,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.customers += 1
	var sub float64
	for i := 0; i < len(product); i++ {
		sub += float64(this.priceDict[product[i]] * amount[i])
	}
	if this.customers%this.lucky == 0 {
		sub = sub * (100 - this.discount) / 100
	}
	return sub
}
