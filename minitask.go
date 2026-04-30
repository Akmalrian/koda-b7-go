package main

import (
	"fmt"

	"github.com/akmalrian/koda7-go/internals/minitask1"
	"github.com/akmalrian/koda7-go/internals/minitask10"
	"github.com/akmalrian/koda7-go/internals/minitask11"
	"github.com/akmalrian/koda7-go/internals/minitask12"
	"github.com/akmalrian/koda7-go/internals/minitask2"
	"github.com/akmalrian/koda7-go/internals/minitask3"
	"github.com/akmalrian/koda7-go/internals/minitask4"
	"github.com/akmalrian/koda7-go/internals/minitask6"
	"github.com/akmalrian/koda7-go/internals/minitask7"
	"github.com/akmalrian/koda7-go/internals/minitask8"
	"github.com/akmalrian/koda7-go/internals/minitask9"
)

func main() {
	for i := true; i == true; {
		fmt.Println("______________________")
		fmt.Println("")
		fmt.Println("Silahkan Pilih Menu : ")
		fmt.Println("1. Minitask 1")
		fmt.Println("2. Minitask 2")
		fmt.Println("3. Minitask 3")
		fmt.Println("4. Minitask 4")
		fmt.Println("5. Minitask 6")
		fmt.Println("6. Minitask 7")
		fmt.Println("7. Minitask 8")
		fmt.Println("8. Minitask 9")
		fmt.Println("9. Minitask 10")
		fmt.Println("10. Minitask 11")
		fmt.Println("11. Minitask 12")
		fmt.Println("0. Keluar")
		fmt.Println()

		fmt.Print("Masukkan Input : ")

		var choice int
		fmt.Scan(&choice)
		fmt.Println()
		if choice == 1 {
			fmt.Println("Pilih Menu : ")
			fmt.Println("1. Hitung Luas Lingkaran")
			fmt.Println("2. Hitung Keliling Lingkaran")
			fmt.Print("Masukkan Input : ")
			var choiceCircle int
			fmt.Scan(&choiceCircle)
			if choiceCircle == 1 {
				fmt.Println("______________________")
				fmt.Println("Hitung Luas Lingkaran ")
				var num uint
				fmt.Print("Masukkan Input : ")
				fmt.Scan(&num)
				luas := minitask1.LuasLingkaran(num)
				fmt.Printf("Luas: %.2f\n", luas)

			} else if choiceCircle == 2 {
				fmt.Println("______________________")
				fmt.Println("Hitung Keliling Lingkaran ")
				var num uint
				fmt.Print("Masukkan Input : ")
				fmt.Scan(&num)
				keliling := minitask1.KelilingLingkaran(num)
				fmt.Printf("Keliling: %.2f\n", keliling)
			} else {
				fmt.Println("Pilihan tidak ada")
			}
		} else if choice == 2 {
			fmt.Println("____________________")
			fmt.Println("Cetak Segitiga")
			fmt.Print("Masukkan Jumlah Baris : ")
			var num int
			fmt.Scan(&num)
			minitask2.Triangle(num)
		} else if choice == 3 {
			fmt.Println("____________________")
			fmt.Println("Input Slice")
			fmt.Print("Masukkan Angka yang dirubah : ")
			var num int
			fmt.Scan(&num)
			minitask3.SliceInterger([]int{50, 75, 66, 20, 32, 90}, num)
		} else if choice == 4 {
			akmal := minitask4.User{
				Name:   "Akmal Oktarian",
				Photo:  "image.jpg",
				Email:  "akmal@gmail.com",
				Age:    25,
				Phone:  "089603886150",
				Status: true,
				Education: []minitask4.Education{
					{
						Nama:    "UIGM",
						Jurusan: "Sistem Komputer",
					},
					{
						Nama:    "SMA NEGERI 1",
						Jurusan: "IPA",
					},
				},
			}
			fmt.Println((akmal))
		} else if choice == 5 {
			minitask6.FilePath()
		} else if choice == 6 {
			akmal := minitask7.User{
				Name:    "Akmal Oktarian",
				Address: "akmal@gmail.com",
				Phone:   "089603886150",
			}
			akmal.Print()
			fmt.Println()
			fmt.Println((akmal.Greet()))
			fmt.Println()
			fmt.Print("Masukkan Nama yang ingin diganti : ")
			var name string
			fmt.Scan(&name)
			akmal.SetName(name)
			fmt.Println((akmal.Greet()))
		} else if choice == 7 {
			var historyFiktif []int
			var jml int

			fmt.Print("Berapa banyak barang yang ingin dibayar? ")
			fmt.Scan(&jml)

			prices := make([]int, jml)
			for i := 0; i < jml; i++ {
				fmt.Printf("Masukkan harga barang ke-%d: ", i+1)
				fmt.Scan(&prices[i])
			}

			bank := minitask8.Bank{}
			online := minitask8.Online{}
			fiktif := minitask8.Fiktif{History: &historyFiktif}

			fmt.Println("\n--- Hasil Eksekusi ---")

			minitask8.ProcessPayment(bank, prices)
			minitask8.ProcessPayment(online, prices)

			minitask8.ProcessPayment(fiktif, prices)

			fmt.Println("\nData History Pembayaran Fiktif:")
			fmt.Println(historyFiktif)
		} else if choice == 8 {
			minitask9.Coffee()
			// var wg sync.WaitGroup

			// for i := 1; i <= 3; i++ {
			// 	var menu string
			// 	fmt.Scan(&menu)
			// 	go minitask9.Coffee(&wg, menu, i)
			// }
			// wg.Done()

			// wg.Go(minitask9.Coffee)
			// wg.Go(func(){
			// })
		} else if choice == 9 {
			minitask10.Aktivitas()

		} else if choice == 10 {
			minitask11.Message()
		} else if choice == 11 {
			minitask12.Data()
		} else if choice == 0 {
			i = false
			fmt.Println("Telah Keluar")
		} else {
			fmt.Println("Silahkan Pilih Menu yang Lain")
		}
	}
}
