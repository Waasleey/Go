package main

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

const tentativas = 4
const delay = 3

func main() {

	mensageIndrouction()

	for {
		showMenu()

		input := selectMenu()

		switch input {
		case 1:
			fmt.Println("")
			fmt.Println("Init Monitiring")
			fmt.Println("")
			monitoringInit()
		case 2:
			fmt.Println("")
			fmt.Println("logs")
			fmt.Println("")
			exibeLogs()
		case 3:
			fmt.Println("")
			fmt.Println("Exit program")
			fmt.Println("")
			os.Exit(0)
		default:
			fmt.Println("opção invalida")
			os.Exit(-1)
		}
	}

}

func mensageIndrouction() {

	name := "Wasley!"
	years := 22
	versionSystem := 1.1

	fmt.Println("Hello World", name, "You have", years, "old. You are using the system in version", versionSystem)
	fmt.Println("")
}

func showMenu() {

	fmt.Println("1) Init monitoring")
	fmt.Println("2) Show Logs")
	fmt.Println("3) Close the program")
	fmt.Println("")
}

func selectMenu() int {

	var selected int

	fmt.Print("Select the option: ")
	fmt.Scan(&selected)
	return selected
}

func monitoringInit() {

	URLs := leArquivoTxt()

	for i := 0; i <= tentativas; i++ {

		for _, listadesites := range URLs {

			fmt.Println("testando a URL:", listadesites)
			pingSite(listadesites)
		}
		time.Sleep(delay * time.Second)
	}
}

func pingSite(pingURL string) {

	resp, err := http.Get(pingURL)

	if err != nil {
		fmt.Println("Ocorreu um erro durante a execução:", err)
		fmt.Println("")
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", pingURL, "está sem problemas")
		registraLog(pingURL, true)
		fmt.Println("")
	} else {
		fmt.Println("o site:", pingURL, "está com problemas, o status HTTP do site é", resp.StatusCode)
		registraLog(pingURL, false)
		fmt.Println("")
	}
}

func leArquivoTxt() []string {

	var sites []string
	resp, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
	}

	leitura := bufio.NewReader(resp)

	for {

		linha, err := leitura.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	resp.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("Log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo")
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " URL: " + site + " - está com o status: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func exibeLogs() {

	arquivo, err := ioutil.ReadFile("Log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo Logs")
	}

	fmt.Println(string(arquivo))
}
