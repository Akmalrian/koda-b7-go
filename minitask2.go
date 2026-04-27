package main

import "fmt"

func main() {
	square(5)
}

func square(num int) {
	fmt.Println("Hasil Segitiga Siku-Siku ", num, " baris")
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
