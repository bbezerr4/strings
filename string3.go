package main

import (
	"fmt"
	"strings"
)

//Peça ao usuário para digitar uma string e um caractere e retorne o número de
//ocorrências desse caractere na string.

func main() {
	var palavra string
	var caractere string
	fmt.Println("Solte o verbo: ")
	fmt.Scan(&palavra)

	fmt.Println("Digite um caractere: ")
	fmt.Scan(&caractere)

	fmt.Println(strings.Count(palavra, caractere))

}
