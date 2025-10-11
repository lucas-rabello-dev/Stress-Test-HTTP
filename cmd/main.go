package main

import (
	"flag"
	"fmt"

	"github.com/lucas-rabello-dev/Stress-Test-HTTP/internal/model"
)

// CLI aqui

// URL -url
// Num requests -requests
// tempo de duração (timeout)
// flag para salvar dados em um Json (-json <name_file>.json)
// metodo hppt para a request -method

// Default is true
var JsonSave bool = true

func main() {

	url := flag.String("url", "", "URL para as requisições")

	requests := flag.Int("r", 0, "Número de requisições")

	time := flag.Duration("t", 0, "Tempo de execução")

	jsonName := flag.String("json", "", "Nome do arquivo dos dados salvos em Json")

	method := flag.String("m", "", "Método para as requisições HTTP")

	flag.Parse() // inicia as flags

	if *url == "" {
		fmt.Println("Você precisa adicionar um valor em -url")
		return
	}

	if *requests == 0 {
		fmt.Println("Você precisa adicionar um valor em -r")
		return
	}

	if *time == 0 {
		fmt.Println("Você precisa adicionar um valor em -t")
		return
	}

	if *jsonName == "" {
		fmt.Println("Os dados não serão salvos em um Json!")
		JsonSave = false
	}

	if *method == "" {
		fmt.Println("Você precisa adicionar um valor em -m")
		return
	}

	model.addData(url, requests, time, jsonName, method)


}