package main

import "fmt"

func main() {
	n := 7

	// Chama a função isPrime e imprime a mensagem retornada
	_, msg := isPrime(n)
	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 e 1 não são números primos por definição
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// Números negativos não são primos
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// Usa o operador de módulo repetidamente para verificar se é um número primo
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// Não é um número primo
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	// É um número primo
	return true, fmt.Sprintf("%d is a prime number!", n)
}
