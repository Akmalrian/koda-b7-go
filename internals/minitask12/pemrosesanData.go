package minitask12

import (
	"fmt"
	"sync"
	"time"
)

func GenerateNumbers(wg *sync.WaitGroup, limit int, out chan<- int) {
	defer wg.Done()
	defer close(out)
	fmt.Println("--- Tahap 1: Memulai Generate Angka ---")
	for i := 1; i <= limit; i++ {
		time.Sleep(200 * time.Millisecond)
		out <- i
	}
}

func FilterEvenNumbers(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()
	defer close(out)
	fmt.Println("--- Tahap 2: Memulai Filter Angka Genap ---")
	for v := range in {
		time.Sleep(200 * time.Millisecond)
		if v%2 == 0 {
			out <- v
		}
	}
}

func SquareNumbers(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()
	defer close(out)
	fmt.Println("--- Tahap 3: Memulai Hitung Kuadrat ---")
	for v := range in {
		time.Sleep(200 * time.Millisecond)
		out <- v * v
	}
}

func Data() {
	var wg sync.WaitGroup

	fmt.Print("Masukkan Jumlah Angka : ")
	var num int
	fmt.Scan(&num)

	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	wg.Add(4)

	go GenerateNumbers(&wg, num, chan1)
	go FilterEvenNumbers(&wg, chan1, chan2)
	go SquareNumbers(&wg, chan2, chan3)

	go func() {
		defer wg.Done()
		for hasil := range chan3 {
			fmt.Println("Hasil Kuadrat Diterima:", hasil)
		}
	}()
	wg.Wait()
	fmt.Println("\nSemua proses pipeline selesai.")
}
