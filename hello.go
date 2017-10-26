package main

//import do println
import "fmt"

//import "reflect"

//função principal do meu programa
func main() {
	//Sei que veio de um pacote externo pois a função começa com
	// letra maiscula ...

	nome := "Pedro Henrique"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	//fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))

	fmt.Println("1 - Iniciar monitaramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int

	//%d = espera receber um modificador, vai salvar
	// na variavel comando o que for digitado
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço esse comando...")
	}

}
