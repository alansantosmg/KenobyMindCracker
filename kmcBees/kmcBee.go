package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

// stringToFloat converte entrada de stdin para float64
// retorna s(float) e nil se conversao ok
// retorna s(NaN) e mensagem de erro
func stringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(math.NaN()), errors.New("Não é numero")
	}

	return f, nil

}

func main() {

	//imprime instruções de uso
	fmt.Printf(`

kmcBee - Calcula prazo ou quantidade de recursos para uma tarefa

Exemplo 1:
	3000 abelhas --> 50 dias
	4000 abelhas -->  x dias

-->	Execute: kmcBee 3000 50 6000 x

Exemplo 2:
	3000 abelhas --> 50 dias
	x    abelhas --> 37.5 dias

-->	Execute: kmcBee 3000 50 x 37.5

`)
	// verifica quantidade de parametros necessária
	if len(os.Args) == 5 {

		// storage dos parametros de linha de comando
		myArray := make([]float64, 4)

		// percorre range e guarda parametros em c
		for i, arg := range os.Args[1:] {
			c, _ := stringToFloat(arg)
			myArray[i] = c
		}

		// se parametro tempo for desconhecido calcula tempo
		if math.IsNaN(myArray[3]) {
			result := (myArray[0] / myArray[2]) * myArray[1]
			fmt.Println("Previsão de tempo x = ", result)
		}
		// se recursos for desconhecido calcula recursos
		if math.IsNaN(myArray[2]) {
			result := (myArray[0] * myArray[1]) / myArray[3]
			fmt.Println("x recursos necessários = ", result)
		}

	} else {
		// Avisa usuário caso parametros mais ou menos que necessario
		fmt.Println("ATENÇÃO: Número incorreto de parametros. Leia as instruções.")

	}
}
