package main

import (
	"fmt"

	shoppingbasket "jnelle-shopping.basket/shopping_basket"
)

func basketOne() {
	BasketItemsOne := new(shoppingbasket.MyBasket)

	fmt.Printf("\n###########Buy 1 get 1 free###########\n")

	// Buy 1 get 1 free
	// Produkt A0001 hinzufügen, um das zweite Produkt Gratis zu erhalten

	BasketItemsOne.Scan("A0001")
	BasketItemsOne.Scan("A0001")

	// Anzeigen des vollständigen Warenkorbs
	BasketItemsOne.GetProducts()

	// Anzeigen des neuen Gesamtpreises
	BasketItemsOne.Total()

	// Kaufe ein Produkt und kriege das Zweite Gratis
	BasketItemsOne.BuyOneGetOneFree("A0001")

	// Anzeigen des neuen Gesamtpreises
	BasketItemsOne.Total()

	// Anzeigen des vollständigen Warenkorbs
	BasketItemsOne.GetProducts()

}

func basketTwo() {
	BasketItemsTwo := new(shoppingbasket.MyBasket)

	// Produkt A0001 hinzufügen
	BasketItemsTwo.Scan("A0001")

	// Produkt A0002 hinzufügen
	BasketItemsTwo.Scan("A0002")

	// Produkt A0003 hinzufügen
	BasketItemsTwo.Scan("A0003")

	// Gesamtpreis ausgeben
	BasketItemsTwo.Total()

	// Anzeigen des vollständigen Warenkorbs
	BasketItemsTwo.GetProducts()

	// Reduzieren eines Produkts
	BasketItemsTwo.DiscountProduct(0.2, "A0001")

	// Anzeigen des neuen Gesamtpreises
	BasketItemsTwo.Total()

	// Warenkorb leeren
	BasketItemsTwo.RemoveAllProducts()
}

func main() {
	basketOne()
	basketTwo()
}
