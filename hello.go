package main

import (
	"fmt"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Comando não reconhecido")
	}
}

func exibeIntroducao() {
	nome := "Gabriel"
	fmt.Println("Olá, sr", nome)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando escolhido foi", comandoLido)

	return comandoLido
}
