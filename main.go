package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var titular string = "Matheus"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.50

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

}
