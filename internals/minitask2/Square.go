package minitask2

import "fmt"

func Triangle(num int) {
	fmt.Println("Hasil Segitiga Siku-Siku ", num, " baris")
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
