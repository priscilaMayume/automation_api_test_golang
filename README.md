# automation_api_test_golang
<br> <br> 

Este projeto tem como objetivo praticar testes unitários usando a linguagem de programação Go (Golang). Utiliza o ambiente de desenvolvimento Visual Studio Code (VSCode) e o Go na versão go1.22.3 para plataforma Darwin/amd64.
<br> 
Descrição do Projeto
<br> 
O projeto inclui um programa em Go que, ao receber um número inteiro como entrada, identifica se o número é primo ou não, exibindo uma mensagem no terminal com o resultado da verificação.
<br> 
Os testes são projetados para garantir que cada parte do programa funcione corretamente sob várias condições, incluindo entradas válidas, inválidas e casos especiais. Eles cobrem:
<br> <br> 
Validação da funcionalidade principal (isPrime).
Comportamento da interface do usuário (prompt e intro).
Manipulação de entradas do usuário (checkNumbers).
Controle de fluxo e terminação do programa (readUserInput).
<br> <br> 
Passos para Gerar Relatório de Cobertura
Executar os testes com cobertura e gerar um arquivo de saída:
go test -coverprofile=coverage.out
<br> 
Visualizar o relatório de cobertura no terminal:
go tool cover -func=coverage.out
<br> 
Visualizar o relatório de cobertura em formato HTML:
go tool cover -html=coverage.out
<br> <br> 
__________________________________
<br> <br> 
Exemplo de Uso
<br> 
Executar os testes com cobertura:
go test -coverprofile=coverage.out
<br> 
Isso executará os testes e criará um arquivo coverage.out contendo os dados de cobertura.
Visualizar o relatório de cobertura no terminal:
<br> 
go tool cover -func=coverage.out
<br> <br> 
Isso exibirá um resumo da cobertura de cada função no seu código.
Visualizar o relatório de cobertura em HTML:
<br> 
go tool cover -html=coverage.out
<br> <br> 
Isso abrirá um navegador da web mostrando um relatório detalhado da cobertura de código em um formato visual.
<br> <br> 
Resumo
<br> 
go test -coverprofile=coverage.out: Gera um arquivo coverage.out com os dados de cobertura.
<br> 
go tool cover -func=coverage.out: Mostra um resumo da cobertura no terminal.
<br> 
go tool cover -html=coverage.out: Abre um relatório de cobertura em formato HTML.
<br> 
Com esses passos, você poderá calcular e visualizar a cobertura de código dos seus testes em Go.
<br> <br> <br> 


-----------------------------------------------------------------------
<br> <br> <br> 
This project aims to practice unit testing using the Go programming language (Golang). It utilizes the Visual Studio Code (VSCode) development environment and Go version go1.22.3 for the Darwin/amd64 platform.
<br> 
Project Description
<br> 
The project includes a Go program that, when given an integer as input, identifies whether the number is prime or not, displaying a message in the terminal with the result of the verification.
<br> 
Tests are designed to ensure that each part of the program works correctly under various conditions, including valid inputs, invalid inputs and special cases. They cover:
<br> 
Validation of core functionality (isPrime).
<br> 
User interface behavior (prompt and intro).
<br> 
Handling user inputs (checkNumbers).
Program flow control and termination (readUserInput).
<br> <br> 
________________________________________
<br> <br> 
Steps to Generate Coverage Report
Run the tests with coverage and generate an output file:
go test -coverprofile=coverage.out

View the coverage report in the terminal:
go tool cover -func=coverage.out

View the coverage report in HTML format:
go tool cover -html=coverage.out

__________________________________
Example of use
<br> 
Run the tests with coverage:
<br> 
go test -coverprofile=coverage.out
<br> 
This will run the tests and create a coverage.out file containing the coverage data.
View the coverage report in the terminal:
<br> 
go tool cover -func=coverage.out
<br> 
This will display a summary of the coverage of each function in your code.
View the coverage report in HTML:
<br> 
go tool cover -html=coverage.out
<br> 
This will open a web browser showing a detailed code coverage report in a visual format.
<br> <br> 
Summary
<br> 
go test -coverprofile=coverage.out: Generates a coverage.out file with the coverage data.
<br> 
go tool cover -func=coverage.out: Displays a coverage summary in the terminal.
<br> 
go tool cover -html=coverage.out: Opens a coverage report in HTML format.
<br> 
With these steps, you can calculate and visualize the code coverage of your Go tests.
<br> 

