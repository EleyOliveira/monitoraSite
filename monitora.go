package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const espera = 5

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		switch defineOpcao() {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opção desconhecida")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	// := inferência - declara uma variável e define seu tipo
	nome := "Fulano"
	versao := 1.1
	fmt.Println("Olá, sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func defineOpcao() int16 {
	var opcao int16
	fmt.Scan(&opcao)
	fmt.Println("A opção escolhida é:", opcao)	
	return opcao
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := abreArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			monitoraSite(site)
		}
		fmt.Println("")
		time.Sleep(espera * time.Second)
	}
}

func monitoraSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println(err.Error())

	}

	if resposta.StatusCode == 200 {
		fmt.Println("O site ", site, "está ativo", "código", resposta.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "está com problemas", "código", resposta.StatusCode)
		registraLog(site, false)
	}
}

func abreArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err.Error())
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 03:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLog() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(arquivo))
}
