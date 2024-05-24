package main

import (
	"io"
	"os"
	"testing"
)

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

func Test_prompt(t *testing.T) {
	// Salva uma cópia do os.Stdout
	oldOut := os.Stdout

	// Cria um pipe de leitura e escrita
	r, w, _ := os.Pipe()

	// Define os.Stdout para nosso pipe de escrita
	os.Stdout = w

	// Chama a função prompt
	prompt()

	// Fecha nosso pipe de escrita
	_ = w.Close()

	// Reseta os.Stdout para o que era antes
	os.Stdout = oldOut

	// Lê a saída da nossa função prompt() do pipe de leitura
	out, _ := io.ReadAll(r)

	// Realiza nosso teste
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}
