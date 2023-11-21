package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const tentativas = 3
const delay = 5

func main() {

	leArquivo()

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

	URLs := leArquivo()

	for i := 0; i < tentativas; i++ {

		for _, sites := range URLs {

			fmt.Println("testando a URL:", sites)
			testSite(sites)
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(sites string) {

	resp, err := http.Get(sites)

	if err != nil {
		fmt.Println("Ocorreu um erro durante a execução:", err)
		fmt.Println("")
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", sites, "está sem problemas")
		fmt.Println("")
	} else {
		fmt.Println("o site:", sites, "está com problemas, o status HTTP do site é", resp.StatusCode)
		fmt.Println("")
	}
}

func leArquivo() []string {

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
