package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_alpha_isPrime(t *testing.T) {
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

func Test_alpha_prompt(t *testing.T) {
	// Salva uma cópia do os.Stdout
	oldOut := os.Stdout

	// Cria um pipe de leitura e escrita
	r, w, _ := os.Pipe()

	// Define os.Stdout para nosso pipe de escrita
	os.Stdout = w

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

func Test_alpha_intro(t *testing.T) {
	// Salva uma cópia do os.Stdout
	oldOut := os.Stdout

	// Cria um pipe de leitura e escrita
	r, w, _ := os.Pipe()

	// Define os.Stdout para nosso pipe de escrita
	os.Stdout = w

	intro()

	// Fecha nosso pipe de escrita
	_ = w.Close()

	// Reseta os.Stdout para o que era antes
	os.Stdout = oldOut

	// Lê a saída da nossa função intro() do pipe de leitura
	out, _ := io.ReadAll(r)

	// Realiza nosso teste
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct, got %s", string(out))
	}
}

func Test_alpha_checkNumbers(t *testing.T) {
	// Define os casos de teste
	test := []struct {
		name     string
		input    string
		expected string
}{
	{name: "empty", input: "", expected: "Please enter a whole number!"},
	{name: "prime_number", input: "7", expected: "7 is a prime number!"},
	{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
	{name: "negative", input: "-5", expected: "Negative numbers are not prime, by definition!"},
	{name: "one", input: "1", expected: "1 is not prime, by definition!"},
	{name: "not_prime", input: "4", expected: "4 is not a prime number because it is divisible by 2"},
	{name: "type", input: "three", expected: "Please enter a whole number!"},
	{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
	{name: "quit", input: "q", expected: ""},
	{name: "Special", input: "@ˆ#*", expected: "Please enter a whole number!"},

}

	for _, e := range test {
		input := strings.NewReader(e.input)
		read := bufio.NewScanner(input)
		res, _ := checkNumbers(read)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected, %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_alpha_readUserInput(t *testing.T) {
	// Para testar esta função, precisamos de um canal e uma instância de um io.Reader
	doneChan := make(chan bool)

	// Cria uma referência a um bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}