package minitask6

import (
	"fmt"
	"io"
	"os"
)

func FilePath() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC : ", r)
			fmt.Println("Proses dilanjutkan...")
		}
	}()

	fmt.Print("Masukkan nama file yang ingin dibuka: ")
	var fileName string
	fmt.Scan(&fileName)

	path := "./internals/" + fileName

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error saat membuka file:", err)
		return
	}

	defer func() {
		fmt.Println("Menutup file...")
		file.Close()
	}()

	fileInfo, _ := file.Stat()
	if fileInfo.IsDir() {
		panic("Path yang dimasukkan adalah direktori, bukan file!")
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Isi File:")
	fmt.Println(string(content))
}
