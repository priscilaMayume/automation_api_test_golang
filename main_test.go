package main

import "testing"

func Test_isPrime(t *testing.T) {
	// Define os casos de teste
	primeTest := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		// Caso de teste: número primo
		{"prime", 7, true, "7 is a prime number!"},
		{"not_prime", 4, false, "4 is not a prime number because it is divisible by 2"},
		{"negative", -5, false, "Negative numbers are not prime, by definition!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
	}

	// Itera sobre cada caso de teste
	for _, e := range primeTest {
		// Chama a função isPrime com o número de teste
		result, msg := isPrime(e.testNum)

		// Verifica se o resultado booleano está correto
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		// Verifica se a mensagem retornada está correta
		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
