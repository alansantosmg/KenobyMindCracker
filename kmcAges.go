package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func stringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		err = errors.New("Argumento não é numero")
	}
	return f, err
}

func calculaResultado(mediaIdades1, quantidadePessoas1, mediaIdades2, quantidadePessoas2 float64) float64 {
	mediaIdades1 *= quantidadePessoas1
	mediaIdades2 *= quantidadePessoas2

	if mediaIdades1 > mediaIdades2 {
		return mediaIdades1 - mediaIdades2
	}
	return mediaIdades2 - mediaIdades1
}

func main() {

	if len(os.Args) != 5 {
		fmt.Printf(`

KmcAges - Resolve problema sobre encontrar a idade de uma pessoa usando a diferença de médias de idade entre 2 grupos.

Exemplo:

A Média das idades de Tibério e Ranulfo 45 anos (2 pessoas).
A média das idades de Tibério, Ranulfo e Josefina é 40 anos (3 pessoas.
Qual a idade de Josefina?

Execute: kmcAges 45 2 40 3
`)

	} else {

		commandArgs := make([]float64, 4)

		for i, arg := range os.Args[1:] {
			c, err := stringToFloat(arg)
			commandArgs[i] = c
			if err != nil {
				fmt.Println(err)
			}
		}

		resultado := calculaResultado(commandArgs[0], commandArgs[1], commandArgs[2], commandArgs[3])
		fmt.Println("x = ", resultado)

	}

}
