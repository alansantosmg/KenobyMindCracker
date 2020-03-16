package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//stringToFloat converts string to float64 number
func stringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)

	// if cannot convert return error message
	if err != nil {
		err = errors.New("Argumento não é numero")
	}
	return f, err
}

// calculaResultado gets  age averages and people groups
// then calc the difference of the 2 groups avarages
func calculaResultado(mediaIdades1, quantidadePessoas1, mediaIdades2, quantidadePessoas2 float64) float64 {
	mediaIdades1 *= quantidadePessoas1
	mediaIdades2 *= quantidadePessoas2

	if mediaIdades1 > mediaIdades2 {
		return mediaIdades1 - mediaIdades2
	}
	return mediaIdades2 - mediaIdades1
}

func main() {

	// check if user provided the correct number of arguments
	if len(os.Args) != 5 {

		// print help
		fmt.Printf(`

KmcAges - Resolve problema sobre encontrar a idade de uma pessoa usando a diferença de médias de idade entre 2 grupos.

Exemplo:

A Média das idades de Tibério e Ranulfo 45 anos (2 pessoas).
A média das idades de Tibério, Ranulfo e Josefina é 40 anos (3 pessoas.
Qual a idade de Josefina?

Execute: kmcAges 45 2 40 3
`)

	} else {

		// storage strings converted to floats
		commandArgs := make([]float64, 4)

		//pass os.Args[] to convert values to float
		for i, arg := range os.Args[1:] {
			c, err := stringToFloat(arg)
			commandArgs[i] = c
			if err != nil {
				fmt.Println(err)
			}
		}

		// Return result to user
		resultado := calculaResultado(commandArgs[0], commandArgs[1], commandArgs[2], commandArgs[3])
		fmt.Println("x = ", resultado)

	}

}
