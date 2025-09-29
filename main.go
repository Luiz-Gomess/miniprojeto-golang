package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Gerenciador de Números")

	numeros := []int{3, 4, 5, 6, 7, 8, 9, 2, 143, 56, 32}
	var action string

	MostrarCabecalho()
	fmt.Scanln(&action)

	for action != "0" {

		numeros = Actions(action, numeros)

		MostrarCabecalho()
		fmt.Scanln(&action)
	}

	fmt.Println("Encerrando Aplicação...")
}

func MostrarCabecalho() {
	fmt.Println("+-----------------------------+")
	fmt.Println("|      Escolha sua ação ")
	fmt.Println("|")
	MostrarOpcoes()
	fmt.Print("\nAção: ")
}

func Actions(action string, numeros []int) []int {

	switch action {
	case "1":
		novoNumero, err := NovoNumero()
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			numeros = append(numeros, novoNumero)
			fmt.Println("Numero adicionado: ", novoNumero)
		}

	case "2":
		fmt.Println(numeros)
		time.Sleep(1 * time.Second)

	case "3":
		novosNumeros, err := RemoverPorIndice(numeros)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			numeros = novosNumeros
			time.Sleep(1 * time.Second)
		}

	case "4":
		media, maximo, minimo := Estatisticas(numeros)
		fmt.Println("Média: ", media)
		fmt.Println("Máximo: ", maximo)
		fmt.Println("Mínimo: ", minimo)

	case "5":
		resultado, err := DivisaoSegura()
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("Resultado: ", resultado)
			time.Sleep(1 * time.Second)
		}

	case "6":
		numeros = []int{}
	}

	return numeros
}

func MostrarOpcoes() {
	fmt.Println("| 1 - Adicionar Novo Número")
	fmt.Println("| 2 - Listar Números ")
	fmt.Println("| 3 - Remover Número por índice")
	fmt.Println("| 4 - Estatísticas")
	fmt.Println("| 5 - Divisão Segura")
	fmt.Println("| 6 - Limpar Lista")
	fmt.Println("| 0 - Encerrar Aplicação")
}

// Função para a ação de adicionar um novo Número
func NovoNumero() (int, error) {
	fmt.Println("Digite um número: ")

	var novoNumero int
	fmt.Scanln(&novoNumero)

	if novoNumero < 0 {
		return 0, fmt.Errorf("Número não pode ser negativo!")
	}

	return novoNumero, nil
}

// Função para a ação de remoção
func RemoverPorIndice(numeros []int) ([]int, error) {
	fmt.Println("Digite o número a ser removido: ")

	var indiceDoNumeroARemover int
	fmt.Scanln(&indiceDoNumeroARemover)

	if indiceDoNumeroARemover >= len(numeros) {
		return nil, errors.New("Indice não encontrado.")
	}

	numeros = append(numeros[:indiceDoNumeroARemover], numeros[indiceDoNumeroARemover+1:]...)

	return numeros, nil
}

// Função para a ação de divisão
func DivisaoSegura() (float64, error) {
	fmt.Println("Digite 2 números, primeiro o dividendo e depois o divisor (separados por espaço).")
	var num1, num2 float64
	fmt.Scanln(&num1, &num2)

	if num2 == 0 {
		return 0, errors.New("Não é possível dividir por zero.")
	}
	return num1 / num2, nil
}

// Função para a função de Estatísticas
func Estatisticas(numeros []int) (float64, int, int) {

	if len(numeros) == 0 {
		return 0, 0, 0
	}

	media := Media(numeros)
	maximo := Maximo(numeros)
	minimo := Minimo(numeros)

	return media, maximo, minimo
}

// Funções auxiliares que compõe a função de Estatísticas
func Media(numeros []int) float64 {

	var soma int

	for _, numero := range numeros {
		soma += numero
	}

	return roundFloat(float64(soma)/float64(len(numeros)), 2)
}

// Função auxiliar para arredondar o resultado da média
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Maximo(numeros []int) int {
	var maximo int
	for _, numero := range numeros {
		if numero > maximo {
			maximo = numero
		}
	}

	return maximo
}

func Minimo(numeros []int) int {
	var minimo int = 9999999 // Valor padrão alto para não interferir na primeira iteração
	for _, numero := range numeros {
		if numero < minimo {
			minimo = numero
		}
	}

	return minimo
}
