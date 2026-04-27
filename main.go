// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/akmalrian/koda7-go/internals/art"
// )

// func main() {
// 	var a string = "dfsf"
// 	var b int
// 	var c float32
// 	fmt.Printf("%s,%d,%.5f\n, ", a, b, c)
// 	fmt.Println("hello world", a)
// 	fmt.Println(getUNTotal(70, 75, 82, 79))
// 	total, rata := getUNTotal(75, 87, 92, 81)
// 	fmt.Printf("total = %d, rata-rata = %.2f\n", total, rata)
// 	art.PrintToN(5)

// 	var ages [5]int = [5]int{15, 52, 21, 22, 23}
// 	var scores []int
// 	var ageSlice []int = ages[:]
// 	ageSlice = append(ageSlice, 30)
// 	ageSlice[0] = 10
// 	fmt.Println(ages)
// 	fmt.Println(scores == nil)
// 	fmt.Println(ageSlice)

// 	talents := make([]string, 0, 9)
// 	talents = append(talents, "Akmal", "Aqil", "tes")
// 	fmt.Println(talents)
// 	fmt.Println("Panjang Array = ", len(talents))
// 	fmt.Println("Kapasitasnya = ", cap(talents))
// 	multiHello(talents...)

// 	colors := make(map[string]string)
// 	colors["blue"] = "Biru"
// 	colors["red"] = "Merah"

// 	fmt.Println(colors)
// 	fmt.Println(colors["blue"])
// 	numToDay := map[int]string{
// 		1: "Senin",
// 		2: "Selasa",
// 		3: "Rabu",
// 		4: "Kamis",
// 	}
// 	fmt.Println(numToDay)
// 	fmt.Println(numToDay[3])

// 	dob, _ := time.Parse(time.DateOnly, "2000-10-15")

// 	akmal := user{
// 		id:         1,
// 		name:       "akmal",
// 		email:      "akmal@gmail.com",
// 		isVerified: true,
// 		dob:        dob,
// 	}
// 	fmt.Println((akmal.email))
// 	// scanner := bufio.NewScanner(os.Stdin)
// 	// for scanner.Scan() {
// 	// 	fmt.Println(scanner.Text()) // Membaca baris baru
// 	// }
// }

// type user struct {
// 	id         int
// 	name       string
// 	email      string
// 	isVerified bool
// 	dob        time.Time
// }

// func getUNTotal(mtk uint, eng uint, ind uint, ipa uint) (uint, float32) {
// 	sum := mtk + eng + ind + ipa
// 	avg := float32(sum) / float32(4)

// 	return sum, avg
// }

// func multiHello(names ...string) {
// 	if len(names) == 0 {
// 		return
// 	}
// 	for _, name := range names {
// 		fmt.Printf("Hello %s\n", name)

// 	}
// }
