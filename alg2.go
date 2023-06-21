//Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	vowelCount := countVowels(input)

	fmt.Printf("A string contém %d vogal(is).\n", vowelCount)
}

// Função para contar a quantidade de vogais em uma string
func countVowels(str string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
