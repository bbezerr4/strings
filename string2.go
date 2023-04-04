package main

import "fmt"

//Peça ao usuário para digitar uma string e retorne o número de caracteres nessa string.

func main() {
	var palavra string

	fmt.Println("O queres dizer: ")
	fmt.Scan(&palavra)

	fmt.Print("A contagem de caracteres é: ", len(palavra))

}
