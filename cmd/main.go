package main

import (
	"concorrencia/internal/fetcher"
	"concorrencia/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done
	fmt.Printf("\nTempo total: %s", time.Since(start))
}
