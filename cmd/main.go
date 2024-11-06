package main

import (
	"concorrencia/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		var totalPrice float64
		countPrice := 0.0
		for price := range priceChannel {
			totalPrice += price
			countPrice++
			avgPrice := totalPrice / countPrice
			fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: %.2f \n", price, avgPrice)
		}
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPricesFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPricesFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPricesFromSite3()
	}()

	wg.Wait()
	close(priceChannel)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
