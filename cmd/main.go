package main

import (
	"fmt"

	"github.com/J-yriz/simple-calculator/internal/calculator"
	"github.com/J-yriz/simple-calculator/pkg"
)

func main() {

	task:
		for {
			choise := pkg.InputSelect("Select operation", []string{"Kalkulator Biasa", "Faktorial", "Keluar"})

			switch choise {
			case "Kalkulator Biasa":
				fmt.Printf("Kalkulator Biasa\n[+] Tambah\n[-] Kurang\n[*] Kali\n[/] Bagi\n")
				inputUser := pkg.InputText("Masukan angka", "math")

				result, err := calculator.Calculate(inputUser)
				if err != nil {
					panic(fmt.Sprintf("Failed to calculate %v\n", err))
				}

				fmt.Printf("Hasil: %v\n", result)
			case "Faktorial":
				fmt.Println("Faktorial")
			case "Keluar":
				break task
			}
		}

}
