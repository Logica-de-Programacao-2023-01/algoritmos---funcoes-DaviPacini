//Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do primeiro elemento igual ao
//valor no slice. Caso não encontre, retorne -1.

package main

import "fmt"

func main() {
	numbers := []int{10, 5, 8, 20, 3}
	value := 8
	position := findPosition(numbers, value)
	fmt.Println(position) // Saída: 2
}

// Função para encontrar a posição do primeiro elemento igual ao valor no slice de inteiros
func findPosition(numbers []int, value int) int {
	for i, num := range numbers {
		if num == value {
			return i
		}
	}

	return -1
}
