package main
//import do println
import "fmt"
//import "reflect"

//função principal do meu programa
func main(){
	//Sei que veio de um pacote externo pois a função começa com 
	// letra maiscula ...

	nome := "Pedro Henrique"
	idade := 25
	versao := 1.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade )
	fmt.Println("Este programa está na versão", versao)

	//fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}

