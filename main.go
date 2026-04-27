package main

import "fmt"

func main() {
	var a string = "dfsf"
	var b int
	var c float32
	fmt.Printf("%s,%d,%.5f\n, %d", a, b, c)
	fmt.Println("hello world", a)
	fmt.Println(getUNTotal(70, 75, 82, 79))
	total, rata := getUNTotal(75, 87, 92, 81)
	fmt.Printf("total = %d, rata-rata = %.2f\n", total, rata)
	printToN(5)
}

func getUNTotal(mtk uint, eng uint, ind uint, ipa uint) (uint, float32) {
	sum := mtk + eng + ind + ipa
	avg := float32(sum) / float32(4)

	return sum, avg
}

func printToN(n int) {
	for i := range n {
		fmt.Println(i)
	}
}
