package main

import (
	"fmt"

	"github.com/akmalrian/koda7-go/internals/minitask1"
	"github.com/akmalrian/koda7-go/internals/minitask2"
	"github.com/akmalrian/koda7-go/internals/minitask3"
	"github.com/akmalrian/koda7-go/internals/minitask4"
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
		} else if choice == 0 {
			i = false
			fmt.Println("Telah Keluar")
		} else {
			fmt.Println("Silahkan Pilih Menu yang Lain")
		}
	}
}
