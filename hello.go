package main

import (
	"fmt"
)

func main() {
	nome := "vascao" //definição da variavel nao é obrigatoria, nome := "vascao" funciona igual var nome string = "vascao"
	var versao float32 = 1.1

	fmt.Println("fala ", nome)
	fmt.Println("essa é a versão", versao, "do programa")

	fmt.Println("1- iniciar monitoramento")
	fmt.Println("2- exibir logs")
	fmt.Println("0- Sair")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O valor da variável comando é:", comando)

}
