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
}

func LuasLingkaran(jari uint) float32 {
	return math.Pi * float32(jari) * float32(jari)
}

func KelilingLingkaran(jari uint) float32 {
	return 2 * math.Pi * float32(jari)
}
