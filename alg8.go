//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com apenas os números pares
//contidos no slice. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func filterEvenNumbers(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}

	result := make([]int, 0)
	for _, num := range slice {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result, nil
}

func main() {
	// Exemplo de uso
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNumbers, err := filterEvenNumbers(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", evenNumbers)
	}
}
