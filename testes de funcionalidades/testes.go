package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	// retorna um ponteiro de string
	// o 120 é um valor default
	// ./programa -Flag1 "alguma coisa"
	// output = alguma coisa
	flag1 := flag.String("Flag1", "usando a biblioteca flag", "descrição da flag aqui")

	flag2 := flag.Int("time", 0, "tempo para executar o programa")


	flag.Parse() // usada para inicar as flags

	for i := 0; i < *flag2; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i+1, *flag1)
	}

}