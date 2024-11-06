package processor

import "fmt"

func ShowPriceAVG(priceChannel <-chan float64, done chan<- bool) {
	var totalPrice float64
	countPrice := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrice++
		avgPrice := totalPrice / countPrice
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: %.2f \n", price, avgPrice)
	}

	done <- true
}
