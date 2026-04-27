package main

import (
	"fmt"
	"math"
)

func main() {
	luas := LuasLingkaran(5)
	keliling := KelilingLingkaran(5)

	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)

	//minitask 2
	square(5)

	//minitask 3
	sliceInterger([]int{50, 75, 66, 20, 32, 90})
}

func LuasLingkaran(jari uint) float32 {
	return math.Pi * float32(jari) * float32(jari)
}

func KelilingLingkaran(jari uint) float32 {
	return 2 * math.Pi * float32(jari)
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

func sliceInterger(ageSlice []int) {
	fmt.Println(ageSlice)
	ageSlice = append(ageSlice[:3], append([]int{88}, ageSlice[3:]...)...)
	fmt.Println(ageSlice)

	for i := 0; i < len(ageSlice); i++ {
		fmt.Println(ageSlice[i])
	}
}
