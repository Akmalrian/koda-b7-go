package minitask11

import (
	"fmt"
	"sync"
	"time"
)

type Pesan struct {
	Pengirim string
	Isi      string
}

func Message() {
	var wg sync.WaitGroup
	chn := make(chan Pesan)

	wg.Add(1)
	go PapanTulis(chn, &wg)

	daftarPesan := []Pesan{
		{Pengirim: "Akmal", Isi: "Semangat!"},
		{Pengirim: "Aqil", Isi: "Hebat sekali."},
		{Pengirim: "Ali", Isi: "Ayo pasti bisa!"},
		{Pengirim: "Angga", Isi: "Mantap jiwa."},
	}

	for _, p := range daftarPesan {
		KirimPesan(chn, p.Pengirim, p.Isi)
		time.Sleep(500 * time.Millisecond)
	}

	close(chn)

	wg.Wait()
	fmt.Println("--- Semua pesan telah ditampilkan, channel ditutup ---")
}

func KirimPesan(ch chan<- Pesan, nama string, isi string) {
	time.Sleep(100 * time.Millisecond)
	msg := Pesan{Pengirim: nama, Isi: isi}
	fmt.Printf("-> %s sedang menulis pesan...\n", nama)
	ch <- msg
}

func PapanTulis(ch <-chan Pesan, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Println("---------------------------")
		fmt.Println("[PAPAN TULIS SELESAI MENAMPILKAN]")
	}()
	fmt.Println("[PAPAN TULIS KELUARGA SIAP]")

	for p := range ch {
		fmt.Println("---------------------------")
		fmt.Printf("Dari   : %s\nPesan  : %s\n", p.Pengirim, p.Isi)
		time.Sleep(2 * time.Second)
	}
}

// package minitask11

// import (
//     "fmt"
//     "sync"
//     "time"
// )

// func Message() {
//     var wg sync.WaitGroup
//     var chn chan string
//     chn = make(chan string, 3)
//     pesan := []string{"\nNama : Akmal\nPesan : Semangat", "\nNama : Aqil\nPesan : Hebat", "\nNama : Ali\nPesan : Ayo bisa", "\nNama : Angga\nPesan : Mantap"}
//     // wg.Wait()
//     // wg.Done()
//     papanTulis := 1
//     for range papanTulis {
//         wg.Add(1)
//         go Channel(chn, &wg)
//     }

//     for _, order := range pesan {
//         chn <- order
//         time.Sleep(1 * time.Millisecond)
//         fmt.Println("Pesan masuk")
//         time.Sleep(1 * time.Millisecond)
//     }
//     close(chn)
//     wg.Wait()
//     fmt.Println("Channel Ditutup")
// }

// func Channel(MessageChn chan string, wg *sync.WaitGroup) {
//     defer func() {
//         fmt.Println("Selesai")
//         wg.Done()
//     }()
//     for pesan := range MessageChn {
//         time.Sleep(1 * time.Second)
//         fmt.Printf("Pesan Masuk Dengan isi : %s\n", pesan)

//     }
// }
