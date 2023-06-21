//Crie uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a função recebida
//em cada elemento do slice e retornar um novo slice com os resultados. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func applyFunction(slice []int, fn func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}

	result := make([]int, len(slice))
	for i, val := range slice {
		result[i] = fn(val)
	}

	return result, nil
}

func main() {
	// Exemplo de uso
	slice := []int{1, 2, 3, 4, 5}

	// Função de exemplo para dobrar um número
	double := func(num int) int {
		return num * 2
	}

	// Aplicando a função no slice
	newSlice, err := applyFunction(slice, double)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novo slice:", newSlice)
	}
}
