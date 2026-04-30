package minitask10

import (
	"fmt"
	"sync"
	"time"
)

func Mandi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Mandi")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Mandi")
}
func BuatKopi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Buat Kopi")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Buat Kopi")
}
func Sarapan(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Menyiapkan Sarapan")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Menyiapkan Sarapan")
}
func Kamar(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Merapikan Kamar Tidur")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Merapikan Kamar Tidur")
}

func Aktivitas() {
	var wg sync.WaitGroup
	wg.Add(4)
	go Mandi(&wg)
	go BuatKopi(&wg)
	go Sarapan(&wg)
	go Kamar(&wg)

	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("Berangkat Kerja")
}
