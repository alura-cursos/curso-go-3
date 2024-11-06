package processor

import "fmt"

func ShowPriceAndAVG(priceChannel <-chan float64, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		// fmt.Println(len(priceChannel))
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: R$  %.2f \n", price, avgPrice)
	}
	done <- true
}
