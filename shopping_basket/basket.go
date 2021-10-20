package shoppingbasket

import "math/big"

type basket struct {
	ProductName string
	Price       []*big.Float
}

type MyBasket struct {
	itemList []*basket
}

func getPriceList() []*basket {
	return []*basket{
		{"A0001", []*big.Float{big.NewFloat(12.99)}},
		{"A0002", []*big.Float{big.NewFloat(3.99)}},
		{"A0003", []*big.Float{big.NewFloat(5.99)}},
	}
}
