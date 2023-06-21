//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.

package main

import "fmt"

func main() {
	strings := []string{"Olá", "mundo", "!"}
	concatenated := concatenateStrings(strings)
	fmt.Println(concatenated) // Saída: "Olámundo!"
}

// Função para concatenar strings de um slice
func concatenateStrings(strings []string) string {
	result := ""

	for _, str := range strings {
		result += str
	}

	return result
}
