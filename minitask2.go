package main

import "fmt"

func main() {
	square(5)
}

func square(angkaSquare int) {
	fmt.Println("Hasil Segitiga Siku-Siku ", angkaSquare, " baris")
	for ulang := 1; ulang <= angkaSquare; ulang++ {
		for angka2 := 1; angka2 <= ulang; angka2++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
