package jsonSave

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucas-rabello-dev/Stress-Test-HTTP/internal/model"
)


type Result struct {
	URL string `json:"URL"`
	Requests int `json:"Requests"`
	Time int `json:"Time"` // em model é time.Duration
	JsonFileName string `json:"JsonFileName"`
}

// receber os dados diretamente da struct
func SaveJson(data *model.DataFlag) error {

	result := Result{
		URL: data.URL,
		Requests: data.Requests,
		Time: int(data.Time),
		JsonFileName: data.JsonFileName,
	}

	dadosJson, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return err
	}

	subpasta := "../internal/jsonSave/DataJson"
	caminhoArquivo := filepath.Join(subpasta, fmt.Sprintf("%s.json", data.JsonFileName))

	// cria a subpasta se não existir
	err = os.MkdirAll(subpasta, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(caminhoArquivo)
	if err != nil {
		return err
	}

	_, err = file.Write(dadosJson)
	if err != nil {
		return err
	}

	fmt.Println("Arquivo Json Criado com Sucesso!")

	defer file.Close()
	return nil

}