// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/akmalrian/koda7-go/models"
// 	"github.com/akmalrian/koda7-go/pointer"
// )

// func main() {
// 	fmt.Println("Selamat Datang")
// 	defer func() {
// 		fmt.Println("Sampai Jumpa 1")
// 	}()
// 	defer func() {
// 		fmt.Println("Sampai Jumpa 2")
// 	}()
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("PANIC WOyyy: ", err)

// 		}
// 	}()
// 	// var items = make([]int, 5)
// 	// items[5] = 10

// 	var a string = "dfsf"
// 	var b int
// 	var c float32
// 	fmt.Printf("%s,%d,%.5f\n, ", a, b, c)
// 	fmt.Println("hello world", a)
// 	fmt.Println(getUNTotal(70, 75, 82, 79))
// 	total, rata := getUNTotal(75, 87, 92, 81)
// 	fmt.Printf("total = %d, rata-rata = %.2f\n", total, rata)
// 	if rata < 70 {
// 		os.Exit(0)
// 	}

// 	printToN(5)

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
// 	fmt.Println((akmal.GetEmail()))

// 	pointer.Point()

// 	bulba := models.Bulbasaur{Atk: 10}
// 	pika := models.Pikachu{Atk: 10}
// 	fmt.Println(models.PokemonAtk(bulba))
// 	fmt.Println(models.PokemonAtk(pika))
// 	// models.PokemonAtk(char)
// }

// type user struct {
// 	id         int
// 	name       string
// 	email      string
// 	isVerified bool
// 	dob        time.Time
// }

// // func owner namaFUngsi (parameter) {implementasi}
// func (u user) GetEmail() string {
// 	return u.email
// }

// func getUNTotal(mtk uint, eng uint, ind uint, ipa uint) (uint, float32) {
// 	sum := mtk + eng + ind + ipa
// 	avg := float32(sum) / float32(4)
// 	defer func() {
// 		fmt.Println("Perhitungan Nilai MTK")
// 	}()
// 	return sum, avg
// }

// func printToN(n int) {
// 	for i := range n {
// 		fmt.Println(i)
// 	}
// }

// func multiHello(names ...string) {
// 	if len(names) == 0 {
// 		return
// 	}
// 	for _, name := range names {
// 		fmt.Printf("Hello %s\n", name)

// 	}
// }
