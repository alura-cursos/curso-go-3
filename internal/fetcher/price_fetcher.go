package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

// Buscando preços de diferentes sites
func FetchPrices(priceChannel chan<- float64) {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		defer wg.Done()
		FetchMultiplePriceFromSite4(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPriceFromSite1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFromSite2() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFromSite3() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}

func FetchMultiplePriceFromSite4(priceChannel chan<- float64) {
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}
	for _, price := range prices {
		priceChannel <- price
	}
}
