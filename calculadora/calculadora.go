package main

import (
	"fmt"
	"os"
)

func menu() int {
	op := 6

	for op > 5 {
		fmt.Println("1 - Adição")
		fmt.Println("2 - Subtração")
		fmt.Println("3 - Multiplicação")
		fmt.Println("4 - Divisão")
		fmt.Println("5 - Sair")
		fmt.Scanln(&op)
	}

	return op

}

func adicao(n1 float64, n2 float64) {
	fmt.Printf("A soma de %.2f com %.2f é igual a %.2f\n", n1, n2, n1+n2)
	fmt.Printf("%.2f + %.2f = %.2f\n", n1, n2, n1+n2)
}

func subtracao(n1 float64, n2 float64) {
	fmt.Printf("A subtração de %.2f com %.2f é igual a %.2f\n", n1, n2, n1-n2)
	fmt.Printf("%.2f - %.2f = %.2f\n", n1, n2, n1-n2)
}

func multiplicacao(n1 float64, n2 float64) {
	fmt.Printf("A multiplicação de %.2f com %.2f é igual a %.2f\n", n1, n2, n1*n2)
	fmt.Printf("%.2f * %.2f = %.2f\n", n1, n2, n1*n2)
}

func divisao(n1 float64, n2 float64) {

	if n2 != 0 {
		fmt.Printf("A divisao de %.2f com %.2f é igual a %.2f\n", n1, n2, n1/n2)
		fmt.Printf("%.2f / %.2f = %.2f\n", n1, n2, n1/n2)
	} else {
		fmt.Println("Divisão por zero!!")
	}

}

func main() {

	opcao := 0

	for true {
		opcao = menu()

		if opcao == 5 {
			os.Exit(0)
		}

		var n1 float64
		var n2 float64

		fmt.Println("Digite o primeiro número: ")
		fmt.Scanln(&n1)
		fmt.Println("Digite o segundo número: ")
		fmt.Scanln(&n2)

		switch opcao {
		case 1:
			adicao(n1, n2)
			break
		case 2:
			subtracao(n1, n2)
			break
		case 3:
			multiplicacao(n1, n2)
			break
		case 4:
			divisao(n1, n2)
			break
		default:
			break
		}
	}
}
