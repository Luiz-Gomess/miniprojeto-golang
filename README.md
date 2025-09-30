# Gerenciador de Números em Go

Este projeto tem como objetivo explorar os conceitos básicos da linguagem Go. A aplicação a ser desenvolvida é um Gerenciador de Números executado em linha de comando, que permite ao usuário adicionar, listar, remover e calcular estatísticas sobre números inteiros. O projeto exercita declaração de variáveis, condicionais, laços, funções, slices e funções com múltiplos retornos.

## Funcionalidades

A aplicação deve apresentar ao usuário um menu com as seguintes opções:

* **1) Adicionar número:** solicita um número inteiro e adiciona ao slice.
* **2) Listar números:** mostra todos os números armazenados no slice.
* **3) Remover por índice:** solicita um índice e remove o elemento correspondente.
* **4) Estatísticas:** calcula mínimo, máximo e média dos números (função com múltiplos retornos).
* **5) Divisão segura:** solicita dois números inteiros e realiza a divisão, retornando erro se o divisor for zero.
* **6) Limpar lista:** esvazia o slice de números.
* **0) Sair:** encerra a aplicação.

## Requisitos técnicos

* O programa deve ser implementado em um único arquivo `main.go`.
* O código deve compilar e executar usando o comando `go run main.go`.
* O slice de inteiros deve ser usado para armazenar os números.
* As funções devem estar devidamente definidas para modularizar o código.
* O tratamento de erro deve seguir o padrão do Go (`if err != nil`).

## Funcionalidades bônus

Implemente se houver tempo:

* Impedir que números negativos sejam adicionados.
* Implementar uma opção de ordenação crescente e decrescente da lista.
* Criar uma opção para exibir apenas números pares.
* Implementar exportação da lista para um arquivo texto.
