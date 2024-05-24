# automation_api_test_golang

Este projeto tem como objetivo praticar testes unitários usando a linguagem de programação Go (Golang). Utiliza o ambiente de desenvolvimento Visual Studio Code (VSCode) e o Go na versão go1.22.3 para plataforma Darwin/amd64.

Descrição do Projeto
O projeto inclui um programa em Go que, ao receber um número inteiro como entrada, identifica se o número é primo ou não, exibindo uma mensagem no terminal com o resultado da verificação.

Exemplo de Uso
Execução Rápida dos Testes (em cache):
go test .
Este comando é útil durante o desenvolvimento para verificar rapidamente os resultados dos testes sem a necessidade de reexecutar todos os testes.

Execução Completa dos Testes (sem cache):
go test -count=1 .
Use este comando após realizar mudanças no código ou nos testes para garantir que tudo está sendo testado novamente, sem usar a versão em cache dos resultados.

Execução dos testes em um pacote Go e exibir informações detalhadas sobre o progresso e os resultados dos testes:
go test -v

Análise da Cobertura de Código
Comandos Úteis para Checagem de Cobertura de Testes Unitários:
Exibir Cobertura no Terminal:
go test -cover
Este comando exibe a cobertura de código no terminal após a execução dos testes, mas não armazena esses resultados em um arquivo externo.

Gerar Perfil de Cobertura em Arquivo Externo:
go test -coverprofile=coverage.out
Além de executar os testes e calcular a cobertura de código, este comando armazena os resultados em um arquivo chamado coverage.out. Esse arquivo contém informações detalhadas sobre a cobertura de código que podem ser usadas posteriormente para análises mais detalhadas.

Gerar Relatório HTML da Cobertura de Código:
csharp
go tool cover -html=coverage.out
Utilize este comando para gerar um relatório visual em HTML da cobertura de código a partir do arquivo coverage.out. O relatório inclui detalhes sobre as linhas de código cobertas pelos testes, facilitando a análise e identificação de áreas que precisam de mais testes.

Este projeto é uma oportunidade de aprimorar suas habilidades em testes unitários em Go e utilizar ferramentas que auxiliam na análise e melhoria da qualidade do código.

-----------------------------------------------------------------------

This project aims to practice unit testing using the Go programming language (Golang). It utilizes the Visual Studio Code (VSCode) development environment and Go version go1.22.3 for the Darwin/amd64 platform.

Project Description
The project includes a Go program that, when given an integer as input, identifies whether the number is prime or not, displaying a message in the terminal with the result of the verification.

Usage Example
Quick Test Execution (cached):
go test .
This command is useful during development to quickly check the test results without needing to rerun all the tests.

Complete Test Execution (without cache):
go test -count=1 .
Use this command after making changes to the code or tests to ensure everything is being tested again, without using the cached version of the results.

Running tests on a Go package and displaying detailed information on the progress and results of the tests:
go test -v

Code Coverage Analysis
Useful Commands for Checking Unit Test Coverage:
Display Coverage in the Terminal:
go test -cover
This command displays the code coverage in the terminal after running the tests but does not store these results in an external file.

Generate Coverage Profile in an External File:
go test -coverprofile=coverage.out
In addition to running the tests and calculating code coverage, this command stores the results in a file called coverage.out. This file contains detailed information about the code coverage that can be used later for more detailed analysis.

Generate HTML Report of Code Coverage:
go tool cover -html=coverage.out
Use this command to generate a visual HTML report of the code coverage from the coverage.out file. The report includes details about the lines of code covered by the tests, making it easier to analyze and identify areas that need more tests.

This project is an opportunity to enhance your skills in unit testing in Go and utilize tools that help in analyzing and improving code quality.