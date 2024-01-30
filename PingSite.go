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

func main() {
	name := "Luiz"
	var version float32 = 1.1

	fmt.Println("Olá", name, "versão do projeto", version)

	readTextFile()
	//logs("site", false)

	for {
		input := userOption() //input vai receber o valor do retorno da func

		switch input {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Exibindo logs...\n")
			showLogs()
			fmt.Println("")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1) //mostra pro sistema operacional que algo deu errado
		}
	}
}

func userOption() int {
	var input int

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Sair")
	fmt.Scan(&input)

	return input
}

func monitoring() {
	fmt.Println("Monitorando...")
	urls := []string{"https://cursos.alura.com.br", "https://translate.google.com.br",
		"https://www.youtube.com/watch?v=2kyNEf9IsBQ", "https://go.dev"}

	for i := 0; i < 3; i++ {
		for i, url := range urls {
			fmt.Println("\n", i, "º site:", url)
			testSite(url)
		}
		time.Sleep(10 * time.Second)
	}
}

func testSite(url string) {
	resp, err := http.Get(url)

	if err != nil { //retorna qualquer erro ocorrido
		fmt.Println("Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site está no ar :)")
		logs(url, true)
	} else {
		fmt.Println("Site está fora do ar :(", resp.StatusCode)
		logs(url, false)
	}
}

func readTextFile() []string {
	var urls []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return urls
}

func logs(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //arquivo, flag, permissão

	if err != nil {
		fmt.Println("Erro:", err)
	}

	file.WriteString("Site:" + site + "---- online:" + strconv.FormatBool(status) + "------" +
		time.Now().Format("02/01/2006 15:04:05") + "\n")

	file.Close()
}

func showLogs() {
	file, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(string(file))
}
