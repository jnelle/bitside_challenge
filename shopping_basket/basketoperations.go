package shoppingbasket

import (
	"fmt"
	"math/big"
)

func (mybasket *MyBasket) Scan(item string) {
	PriceSliceCopy := getPriceList()

	for _, i := range PriceSliceCopy {
		if i.ProductName == item {
			items := basket{i.ProductName, i.Price}
			mybasket.itemList = append(mybasket.itemList, &items)

		}
	}
	fmt.Printf("Aktuelle Anzahl von Artikeln im Warenkorb: %v\n", len(mybasket.itemList))
	PriceSliceCopy = nil
}

func (mybasket *MyBasket) GetProducts() {
	for _, i := range mybasket.itemList {
		fmt.Printf("Artikel: %v, Preis: %.2f\n", i.ProductName, i.Price)
	}
}

func (mybasket *MyBasket) Total() {

	var total *big.Float = new(big.Float)
	for _, product := range mybasket.itemList {
		total.Add(total, product.Price[0])
	}

	fmt.Printf("Gesamtpreis beträgt: %.2f€\n", total)
}

func (mybasket *MyBasket) DiscountProduct(discount float64, productName string) {
	realDiscount := (1.00 - discount)
	for _, product := range mybasket.itemList {
		if product.ProductName == productName {
			newPrice := product.Price[0].Mul(product.Price[0], new(big.Float).SetFloat64(realDiscount))
			fmt.Printf("Der Preis von %v wurde um %v %% reduziert.\nDer neue Preis lautet %.2f€\n", product.ProductName, (discount * 100), newPrice)
		}
	}
}

func (mybasket *MyBasket) BuyOneGetOneFree(productName string) {
	var count int = 0
	for _, product := range mybasket.itemList {
		if product.ProductName == productName {
			count++
			if count > 1 {
				newPrice := product.Price[0].Mul(product.Price[0], new(big.Float).SetFloat64(0.5))
				fmt.Printf("Der Preis von %v wurde um 50%% reduziert.\nDer neue Preis lautet %.2f€\n", product.ProductName, newPrice)
			}
		}
	}
}

func (mybasket *MyBasket) RemoveAllProducts() {
	mybasket.itemList = nil
	fmt.Printf("Alle Artikel wurden aus dem Warenkorb entfernt.\n")
}
