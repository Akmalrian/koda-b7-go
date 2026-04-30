package minitask9

import (
	"fmt"
	"sync"
	"time"
)

func Coffee() {
	var wg sync.WaitGroup
	var chn chan string
	chn = make(chan string, 3)
	coffees := []string{"Kopi A", "Kopi B", "Kopi C", "Kopi D"}
	// wg.Wait()
	// wg.Done()
	baristaCount := 3
	for range baristaCount {
		wg.Add(1)
		go Barista(chn, &wg)
	}

	for _, order := range coffees {
		chn <- order
		fmt.Printf("Order masuk untuk menu %s\n", order)
		time.Sleep(1 * time.Millisecond)
	}
	close(chn)
	wg.Wait()
	fmt.Println("Toko Tutup")
}

func Barista(CoffeeChn chan string, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Barista Pulang")
		wg.Done()
	}()
	for coffee := range CoffeeChn {
		time.Sleep(1 * time.Second)
		fmt.Printf("Start making coffee : %s\n", coffee)
		time.Sleep(1 * time.Second)
		fmt.Printf("Finished making coffee : %s\n", coffee)
	}
}
