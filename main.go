package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoMatheus := ContaCorrente{titular: "Matheus", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDaJennie := ContaCorrente{"Jennifer", 699, 654321, 250.0}

	fmt.Println(contaDoMatheus)
	fmt.Println(contaDaJennie)

	var contaDoLucas *ContaCorrente
	contaDoLucas = new(ContaCorrente)
	contaDoLucas.titular = "Lucas"
	contaDoLucas.numeroAgencia = 612
	contaDoLucas.numeroConta = 321456
	contaDoLucas.saldo = 12.0

	fmt.Println(*contaDoLucas)
}