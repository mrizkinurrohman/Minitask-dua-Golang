package tiga

import "fmt"

func ShowDefer(name string) {
	fmt.Println("1. Mulai menjalankan", name)

	defer fmt.Println("3. Selesai menjalankan", name)

	fmt.Println("2. sedang proses", name)
}