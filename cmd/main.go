package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lucas-rabello-dev/Stress-Test-HTTP/internal/jsonSave"
	"github.com/lucas-rabello-dev/Stress-Test-HTTP/internal/model"
)

// CLI aqui

// URL -url
// Num requests -requests
// tempo de duração (timeout)
// flag para salvar dados em um Json (-json <name_file>.json)

// Default is true
var JsonSave bool = true

func main() {

	url := flag.String("url", "", "URL para as requisições")

	requests := flag.Int("requests", 0, "Número de requisições")

	time_ := flag.Duration("time", 0, "Tempo de execução")

	jsonName := flag.String("json", "", "Nome do arquivo dos dados salvos em Json")

	flag.Parse() // inicia as flags

	if *url == "" {
		fmt.Println("Você precisa adicionar um valor em -url")
		return
	}

	if *requests == 0 {
		fmt.Println("Você precisa adicionar um valor em -r")
		return
	}

	if *time_ == 0 {
		fmt.Println("Você precisa adicionar um valor em -t")
		return
	}

	if *jsonName == "" {
		fmt.Println("Os dados não serão salvos em um Json!")
		JsonSave = false
	}

	instance := model.AddData(url, requests, time_, jsonName)



	if JsonSave {
		err := jsonSave.SaveJson(instance)	
		if err != nil {
			log.Fatal(err)
		}
	}


}