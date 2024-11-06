package processor

import (
	"concorrencia/internal/models"
	"fmt"
)

func ShowPriceAndAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		// fmt.Println(len(priceChannel))
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("[%s]Preço recebido de %s | R$ %.2f | Preço médi até agora %.2f \n", price.TimeStamp.Format("02-Jan 15:04:05"), price.StoreName, price.Value, avgPrice)
	}
	done <- true
}
