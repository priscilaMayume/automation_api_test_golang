# automation_api_test_golang

Este projeto tem como objetivo praticar testes unitários usando a linguagem de programação Go (Golang). Utiliza o ambiente de desenvolvimento Visual Studio Code (VSCode) e o Go na versão go1.22.3 para plataforma Darwin/amd64.

Descrição do Projeto
O projeto inclui um programa em Go que, ao receber um número inteiro como entrada, identifica se o número é primo ou não, exibindo uma mensagem no terminal com o resultado da verificação.

Os testes são projetados para garantir que cada parte do programa funcione corretamente sob várias condições, incluindo entradas válidas, inválidas e casos especiais. Eles cobrem:

Validação da funcionalidade principal (isPrime).
Comportamento da interface do usuário (prompt e intro).
Manipulação de entradas do usuário (checkNumbers).
Controle de fluxo e terminação do programa (readUserInput).

Passos para Gerar Relatório de Cobertura
Executar os testes com cobertura e gerar um arquivo de saída:
go test -coverprofile=coverage.out

Visualizar o relatório de cobertura no terminal:
go tool cover -func=coverage.out

Visualizar o relatório de cobertura em formato HTML:
go tool cover -html=coverage.out

__________________________________

Exemplo de Uso
Executar os testes com cobertura:
go test -coverprofile=coverage.out

Isso executará os testes e criará um arquivo coverage.out contendo os dados de cobertura.
Visualizar o relatório de cobertura no terminal:
go tool cover -func=coverage.out

Isso exibirá um resumo da cobertura de cada função no seu código.
Visualizar o relatório de cobertura em HTML:
go tool cover -html=coverage.out

Isso abrirá um navegador da web mostrando um relatório detalhado da cobertura de código em um formato visual.
Resumo
go test -coverprofile=coverage.out: Gera um arquivo coverage.out com os dados de cobertura.
go tool cover -func=coverage.out: Mostra um resumo da cobertura no terminal.
go tool cover -html=coverage.out: Abre um relatório de cobertura em formato HTML.
Com esses passos, você poderá calcular e visualizar a cobertura de código dos seus testes em Go.


-----------------------------------------------------------------------

This project aims to practice unit testing using the Go programming language (Golang). It utilizes the Visual Studio Code (VSCode) development environment and Go version go1.22.3 for the Darwin/amd64 platform.

Project Description
The project includes a Go program that, when given an integer as input, identifies whether the number is prime or not, displaying a message in the terminal with the result of the verification.

Tests are designed to ensure that each part of the program works correctly under various conditions, including valid inputs, invalid inputs and special cases. They cover:

Validation of core functionality (isPrime).
User interface behavior (prompt and intro).
Handling user inputs (checkNumbers).
Program flow control and termination (readUserInput).

________________________________________

Steps to Generate Coverage Report
Run the tests with coverage and generate an output file:
go test -coverprofile=coverage.out

View the coverage report in the terminal:
go tool cover -func=coverage.out

View the coverage report in HTML format:
go tool cover -html=coverage.out

__________________________________
Example of use
Run the tests with coverage:
go test -coverprofile=coverage.out

This will run the tests and create a coverage.out file containing the coverage data.
View the coverage report in the terminal:
go tool cover -func=coverage.out

This will display a summary of the coverage of each function in your code.
View the coverage report in HTML:
go tool cover -html=coverage.out

This will open a web browser showing a detailed code coverage report in a visual format.
Summary
go test -coverprofile=coverage.out: Generates a coverage.out file with the coverage data.
go tool cover -func=coverage.out: Displays a coverage summary in the terminal.
go tool cover -html=coverage.out: Opens a coverage report in HTML format.
With these steps, you can calculate and visualize the code coverage of your Go tests.

