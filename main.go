package main

import (
	"fmt"
	"if-else/primenumbers"
)

func main() {
	var number int

	fmt.Println("Digite um número:")

	fmt.Scanln(&number)

	isPrime := primenumbers.IsPrimeNumber(number)

	if isPrime {
		fmt.Printf("O número %d é primo", number)
	} else {
		fmt.Printf("O número %d não é primo", number)
	}
}
