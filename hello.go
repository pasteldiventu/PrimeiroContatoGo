package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	for {

		exibeMenu()

		comando := leComando()

		switch comando {

		case 1:
			iniciarMonitoramento()

		case 2:
			fmt.Println("Exibindo logs...")

		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)

		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "vascao" //definição da variavel nao é obrigatoria, nome := "vascao" funciona igual var nome string = "vascao"
	var versao float32 = 1.2

	fmt.Println("fala ", nome)
	fmt.Println("essa é a versão", versao, "do programa")
}

func exibeMenu() {
	fmt.Println("1- iniciar monitoramento")
	fmt.Println("2- exibir logs")
	fmt.Println("0- Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://univagead.com.br/avaunivag"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " está funcionando corretamente!")
	} else {
		fmt.Println("o Site: ", site, "deu erro no carregamento. Status code:", resp.StatusCode)
	}
}
