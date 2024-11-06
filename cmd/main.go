package main

import (
	"concorrencia/internal/fetcher"
	"concorrencia/internal/models"
	"concorrencia/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// buffer
	priceChannel := make(chan models.PriceDetail, 4)
	// len (valores no buffer) cap (tamanho máximo)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAndAVG(priceChannel, done)

	<-done

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
