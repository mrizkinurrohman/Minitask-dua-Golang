package main

import (
	"fmt"
	"minitaskdua/dua"
	"minitaskdua/empat"
	"minitaskdua/lima"
	"minitaskdua/satu"
	"minitaskdua/tiga"
)

func main() {

	// nomor1
	data := []int{1, 2, 3, 4, 5}
	total := satu.SumArray(data)
	fmt.Println("Total value dari elemen adalah", total)

	//nomor2
	nama := "Rizki"
	fmt.Printf("nama saya %s\n", nama)
	dua.ShowPointer(&nama, "Rohman")
	fmt.Printf("nama panggilan %s\n", nama)

	//nomor3
	fmt.Println("Start Defer")
	tiga.ShowDefer("Defer")
	fmt.Println("End Defer")

	//nomor4
	// rohman := empat.Person{
	// 	Name : "Rizki",
	// 	Address: "Bogor",
	// 	Phone: "081234567",
	// }
	// fmt.Println(rohman.GetPersonName())
	// fmt.Println(rohman.GetPersonPhone())
	// // fmt.Println(rohman.GetPersonAddress())
	// rohman.GreetName("Rohman")
	// fmt.Println("Nama Panggilan", rohman.Name)

	rohman := empat.NewPerson("Rohman", "Bogor", "082123456789")
	fmt.Println(rohman.GetPersonName())
	fmt.Println(rohman.GetPersonAddress())
	fmt.Println(rohman.GetPersonPhone())
	rohman.GreetName("rizki")
	fmt.Println("nama berubah menjadi", rohman.GetPersonName())

	//interface
	printJenis(lima.Home{})
	printJenis(lima.Subsidi{})

}

type tipeHome interface {
	Jenis()
}

func printJenis(t tipeHome){
	t.Jenis()
}

