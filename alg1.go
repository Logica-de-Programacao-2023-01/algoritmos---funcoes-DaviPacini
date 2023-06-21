//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

package main

import "fmt"

func main() {
	var (
		n     int
		valor int
		soma  int
	)

	numeros := []int{}
	fmt.Println("Informe quantos valores deseja informar:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Informe um valor")
		fmt.Scan(&valor)
		numeros = append(numeros, valor)
	}
	for z := 0; z < len(numeros); z++ {

		soma += numeros[z]
	}
	media := soma / len(numeros)
	fmt.Println("A média aritmética é:", media)
}
