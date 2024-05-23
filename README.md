# automation_api_test_golang

Este projeto tem como objetivo praticar testes unitários usando a linguagem de programação Go (Golang). Utiliza o ambiente de desenvolvimento Visual Studio Code (VSCode) e o Go na versão go1.22.3 para plataforma Darwin/amd64.

Descrição do Projeto
O projeto inclui um programa em Go que, ao receber um número inteiro como entrada, identifica se o número é primo ou não, exibindo uma mensagem no terminal com o resultado da verificação.

Exemplo de Uso
Execução Rápida dos Testes (em cache):

bash
Copiar código
go test .
Este comando é útil durante o desenvolvimento para verificar rapidamente os resultados dos testes sem a necessidade de reexecutar todos os testes.

Execução Completa dos Testes (sem cache):

bash
Copiar código
go test -count=1 .
Use este comando após realizar mudanças no código ou nos testes para garantir que tudo está sendo testado novamente, sem usar a versão em cache dos resultados.

Análise da Cobertura de Código
Comandos Úteis para Checagem de Cobertura de Testes Unitários:
Exibir Cobertura no Terminal:

bash
Copiar código
go test -cover
Este comando exibe a cobertura de código no terminal após a execução dos testes, mas não armazena esses resultados em um arquivo externo.

Gerar Perfil de Cobertura em Arquivo Externo:

bash
Copiar código
go test -coverprofile=coverage.out
Além de executar os testes e calcular a cobertura de código, este comando armazena os resultados em um arquivo chamado coverage.out. Esse arquivo contém informações detalhadas sobre a cobertura de código que podem ser usadas posteriormente para análises mais detalhadas.

Gerar Relatório HTML da Cobertura de Código:

csharp
Copiar código
go tool cover -html=coverage.out
Utilize este comando para gerar um relatório visual em HTML da cobertura de código a partir do arquivo coverage.out. O relatório inclui detalhes sobre as linhas de código cobertas pelos testes, facilitando a análise e identificação de áreas que precisam de mais testes.

Este projeto é uma oportunidade de aprimorar suas habilidades em testes unitários em Go e utilizar ferramentas que auxiliam na análise e melhoria da qualidade do código.
