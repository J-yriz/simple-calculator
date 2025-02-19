package main

import (
	"fmt"

	"github.com/J-yriz/simple-calculator/internal/calculator"
	"github.com/J-yriz/simple-calculator/pkg"
)

func main() {

task:
	for {
		choise := pkg.InputSelect("Select operation", []string{"Kalkulator Biasa", "Faktorial", "Berpangkat", "Radical", "Keluar"})

		switch choise {
		case "Kalkulator Biasa":
			fmt.Printf("Kalkulator Biasa\n[+] Tambah\n[-] Kurang\n[*] Kali\n[/] Bagi\n")
			inputUser := pkg.InputText("Masukan angka", "math")

			result, err := calculator.Calculate(inputUser)
			if err != nil {
				fmt.Printf("Failed to calculate %v\n", err)
				continue
			}

			fmt.Printf("\nHasil dari %s adalah %v\n", inputUser, result)
		case "Faktorial":
			inputUser := pkg.InputText("Masukan angka", "number")

			// Convert string to int
			number, err := pkg.StrToNmbr(inputUser)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			result := calculator.Factorial(uint32(number))
			fmt.Printf("Hasil dari %s! adalah %d\n", inputUser, result)
		case "Berpangkat":
			inputUser1 := pkg.InputText("Masukan angka", "number")
			inputUser2 := pkg.InputText("Masukan pangkat", "number")

			// Convert string to int
			number, err := pkg.StrToNmbr(inputUser1)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			square, err := pkg.StrToNmbr(inputUser2)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			result := calculator.Square(uint32(number), uint16(square))
			fmt.Printf("\nHasil dari %s pangkat %s adalah %d\n", inputUser1, inputUser2, result)
		case "Radical":
			inputUser := pkg.InputText("Masukan angka", "number")

			// Convert string to int
			number, err := pkg.StrToNmbr(inputUser)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			result := calculator.Radical(uint32(number))
			fmt.Printf("\nHasil dari √%s adalah %.2f\n", inputUser, result)
		case "Keluar":
			break task
		}
	}

}
