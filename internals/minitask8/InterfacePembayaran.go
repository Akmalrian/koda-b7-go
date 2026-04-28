package minitask8

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Payment(prices []int) (int, error)
}

type Bank struct{}

func (b Bank) Payment(prices []int) (int, error) {
	total := 0
	for _, p := range prices {
		if p <= 0 {
			return 0, errors.New("harga <= 0 tidak valid")
		}
		total += p
	}
	fmt.Printf("[BANK] Pembayaran Berhasil! Total: Rp.%d\n", total)
	return total, nil
}

type Online struct{}

func (o Online) Payment(prices []int) (int, error) {
	total := 0
	for _, p := range prices {
		if p <= 0 {
			return 0, errors.New("harga <= 0 tidak valid")
		}
		total += p
	}
	fmt.Printf("[ONLINE] Pembayaran Berhasil! Total: Rp.%d\n", total)
	return total, nil
}

type Fiktif struct {
	History *[]int
}

func (f Fiktif) Payment(prices []int) (int, error) {
	total := 0
	for _, p := range prices {
		if p <= 0 {
			return 0, errors.New("harga <= 0 tidak valid")
		}
		total += p
	}

	*f.History = append(*f.History, total)
	return total, nil
}

func ProcessPayment(pm PaymentMethod, prices []int) {
	_, err := pm.Payment(prices)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
