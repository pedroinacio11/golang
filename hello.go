package main

//import do println
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constantes Não podem ser alteradas

const monitoramentos = 3
const delay = 5

//import "reflect"

//função principal do meu programa
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
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			//Saindo do programa com sucesso (Boa prática o retorno
			//deve ser status zero - estou avisando para o meu SO que sai)
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando...")
			// -1 indica que houve um problema no programa
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	//Sei que veio de um pacote externo pois a função começa com
	// letra maiscula ...

	nome := "Pedro Henrique"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	//fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}

func exibeMenu() {

	fmt.Println("1 - Iniciar monitaramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

//função que retorna um int
func leComando() int {

	var comandoLido int
	//%d = espera receber um modificador, vai salvar
	// na variavel comando o que for digitado
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando Monitorando...")
	//sites := []string{"https://random-status-code.herokuapp.com/", "http://globoesporte.globo.com/", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	// o range obtem a posiçaõ e quem está na posição...
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
	}

	fmt.Println("")
}

func testaSite(site string) {
	//resp, err := http.Get(site)
	//_ serve para ignorar o retorno

	resp, err := http.Get(site)
	//fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O seu site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O seu site", site, "está com algum problema", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		// Tô informando o byte limitador ... ele deve quebrar na linha
		linha, err := leitor.ReadString('\n')
		//TrimSpace = Quebra de espaço -- Não faz ele pular a linha
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		//EOF é o erro que dá quando chega ao final do arquivo
		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)

	//fmt.Println(string(arquivo))

	//Fechando para ser usado por outro processo [boa pratica]
	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}
func imprimeLogs() {

	//ioutil ela abre e fecha sozinha, eu não preciso fechar
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	//convertendo byte para string
	fmt.Println(string(arquivo))

}
