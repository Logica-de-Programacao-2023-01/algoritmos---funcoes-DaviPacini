//Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{10, 5, 8, 20, 3}
	secondLargest := findSecondLargest(numbers)
	fmt.Println(secondLargest) // Saída: 10
}

// Função para encontrar o segundo maior valor em um slice de inteiros
func findSecondLargest(numbers []int) int {
	if len(numbers) < 2 {
		panic("O slice deve conter pelo menos dois elementos.")
	}

	// Ordena o slice em ordem decrescente
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	return numbers[1]
}
