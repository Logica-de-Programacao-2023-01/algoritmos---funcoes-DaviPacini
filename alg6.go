//Escreva uma função que receba um slice de strings como parâmetro e retorne uma string com todas as strings
//concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func main() {
	strings := []string{"Olá", "mundo", "!"}
	concatenated, err := concatenateWithComma(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(concatenated) // Saída: "Olá,mundo,!"
	}

	emptyStrings := []string{}
	emptyConcatenated, err := concatenateWithComma(emptyStrings)
	if err != nil {
		fmt.Println("Erro:", err) // Saída: "Slice vazio"
	} else {
		fmt.Println(emptyConcatenated)
	}
}

// Função para concatenar strings com vírgula
func concatenateWithComma(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("Slice vazio")
	}

	var concatenated = strings.Join()
	return concatenated, nil
}
